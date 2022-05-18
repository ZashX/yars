package dataset

import (
	"context"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"gitlab.com/yars-ai/yars/src"
	"golang.org/x/sync/errgroup"
)

type MF struct {
	rng        src.RandomGenerator
	nFactors   int
	topK       int
	nSamples   int
	lr         float32
	reg        float32
	initMean   float32
	initStdDev float32
}

type User struct {
	Id         string    `db:"user_id"`
	Vec        []float32 `db:"vec"`
	Age        int       `db:"age"`
	Gender     string    `db:"gender"`
	Occupation string    `db:"occupation"`
}

type Item struct {
	Id  string    `db:"item_id"`
	Vec []float32 `db:"vec"`
}

type Interaction struct {
	UserId string `db:"user_id"`
	ItemId string `db:"item_id"`
}

func (mf *MF) InitTableMovielens100k(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector;")
	if err != nil {
		return errors.Wrap(err, "Can't create extension")
	}

	_, err = conn.Exec(ctx, "DROP TABLE IF EXISTS users")
	if err != nil {
		return errors.Wrap(err, "Can't drop table users")
	}
	_, err = conn.Exec(
		ctx,
		fmt.Sprintf(
			`
			CREATE TABLE users (
				user_id varchar(256) NOT NULL,
				age smallint NOT NULL,
				gender varchar(32) NOT NULL,
				occupation varchar(256) NOT NULL,
				vec vector( %d ) NOT NULL,
				PRIMARY KEY(user_id)
			);`, mf.nFactors))
	if err != nil {
		return errors.Wrap(err, "Can't create table users")
	}

	_, err = conn.Exec(ctx, "DROP TABLE IF EXISTS items")
	if err != nil {
		return errors.Wrap(err, "Can't drop table items")
	}
	_, err = conn.Exec(
		ctx,
		fmt.Sprintf(
			`
			CREATE TABLE items (
				item_id varchar(256) NOT NULL,
				vec vector( %d ) NOT NULL,
				PRIMARY KEY(item_id)
			);`,
			mf.nFactors))
	if err != nil {
		return err
	}

	userFile, err := os.Open("./movielens100k/u.user")
	if err != nil {
		return errors.Wrap(err, "Can't open file")
	}
	defer userFile.Close()

	res, err := ioutil.ReadAll(userFile)
	if err != nil {
		return errors.Wrap(err, "Can't read file")
	}
	dataLines := strings.Split(string(res), "\n")

	for _, l := range dataLines {
		vals := strings.Split(l, "|")
		age, err := strconv.Atoi(vals[1])
		if err != nil {
			return errors.Wrap(err, "Can't convert age to int")
		}
		u := User{
			Id:         vals[0],
			Age:        age,
			Gender:     vals[2],
			Occupation: vals[3],
			Vec:        src.NormIf(mf.rng.NewNormalVector(mf.nFactors, mf.initMean, mf.initStdDev)),
		}
		conn.Exec(ctx, "INSERT INTO users VALUES ($1, $2, $3, $4, $5::float4[])", u.Id, u.Age, u.Gender, u.Occupation, u.Vec)
		if err != nil {
			return errors.Wrap(err, "Can't insert user into users")
		}
	}

	itemFile, err := os.Open("./movielens100k/u.item")
	if err != nil {
		return errors.Wrap(err, "Can't open file")
	}
	defer itemFile.Close()

	res, err = ioutil.ReadAll(itemFile)
	if err != nil {
		return errors.Wrap(err, "Can't read file")
	}
	dataLines = strings.Split(string(res), "\n")

	for _, l := range dataLines {
		vals := strings.Split(l, "|")
		i := Item{
			Id:  vals[0],
			Vec: src.NormIf(mf.rng.NewNormalVector(mf.nFactors, mf.initMean, mf.initStdDev)),
		}
		_, err = conn.Exec(ctx, "INSERT INTO items VALUES ($1, $2::float4[])", i.Id, i.Vec)
		if err != nil {
			return errors.Wrap(err, "Can't insert item into items")
		}
	}

	return nil
}

func (mf *MF) Fit(ctx context.Context, conn *pgx.Conn, userId, itemId string) error {
	var user User
	var userSamples []*User
	err := pgxscan.Get(ctx, conn, &user, "SELECT user_id, vec::float4[] FROM users WHERE user_id = $1", userId)
	if err != nil {
		return errors.Wrap(err, "Can't get interact user")
	}

	err = pgxscan.Select(
		ctx,
		conn,
		&userSamples,
		`
		SELECT user_id, vec::float4[]
		FROM users
		WHERE user_id NOT IN (
			SELECT user_id
			FROM train
			WHERE (item_id = $1))
		ORDER BY random() LIMIT $2
		`,
		userId,
		mf.nSamples)
	if err != nil {
		return errors.Wrap(err, "Can't select user samples")
	}

	var item Item
	var itemSamples []*Item
	err = pgxscan.Get(ctx, conn, &item, "SELECT item_id, vec::float4[] FROM items WHERE item_id = $1", itemId)
	if err != nil {
		return errors.Wrap(err, "Can't get interact item")
	}

	err = pgxscan.Select(
		ctx,
		conn,
		&itemSamples,
		`
		SELECT item_id, vec::float4[]
		FROM items
		WHERE item_id NOT IN (
			SELECT item_id
			FROM train
			WHERE (user_id = $1))
		ORDER BY random() LIMIT $2
		`,
		userId,
		mf.nSamples)
	if err != nil {
		return errors.Wrap(err, "Can't select item samples")
	}

	for _, is := range itemSamples {
		sub := src.SubVV(item.Vec, is.Vec)
		diff := src.Dot(user.Vec, sub)
		expDiff := float32(math.Exp(-float64(diff)))
		gradMult := expDiff / (1.0 + expDiff)
		grad := src.MulVS(sub, gradMult*mf.lr)
		user.Vec = src.Norm(src.AddVV(user.Vec, grad))
	}

	for _, us := range userSamples {
		sub := src.SubVV(user.Vec, us.Vec)
		diff := src.Dot(item.Vec, sub)
		expDiff := float32(math.Exp(-float64(diff)))
		gradMult := expDiff / (1.0 + expDiff)
		grad := src.MulVS(sub, gradMult*mf.lr)
		item.Vec = src.Norm(src.AddVV(item.Vec, grad))
	}

	_, err = conn.Exec(ctx, "UPDATE users SET vec = $1::float4[] WHERE user_id = $2", user.Vec, userId)
	if err != nil {
		return errors.Wrap(err, "Can't update user")
	}

	_, err = conn.Exec(ctx, "UPDATE items SET vec = $1::float4[] WHERE item_id = $2", item.Vec, item.Id)
	if err != nil {
		return errors.Wrap(err, "Can't update item")
	}

	return nil
}

func FitTrain(ctx context.Context, pool *pgxpool.Pool, mf *MF) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	var trainInterations []*Interaction
	err = pgxscan.Select(ctx, conn.Conn(), &trainInterations, "SELECT user_id, item_id FROM train ORDER BY random()")
	if err != nil {
		return err
	}
	conn.Release()

	errs, ctx := errgroup.WithContext(ctx)
	for i := range trainInterations {
		userId := trainInterations[i].UserId
		itemId := trainInterations[i].ItemId
		errs.Go(func() error {
			c, e := pool.Acquire(ctx)
			if e != nil {
				return e
			}
			e = mf.Fit(ctx, c.Conn(), userId, itemId)
			defer c.Release()
			if e != nil {
				return e
			}
			return nil
		})
	}
	err = errs.Wait()
	if err != nil {
		return err
	}
	return nil
}

func NDCG(predict []bool, topK int) float32 {
	predict = predict[:topK]
	var iDCG, DCG float32 = 0, 0
	var cnt int = 0
	for i := range predict {
		if predict[i] {
			iDCG += float32(1 / math.Log2(float64(cnt+2)))
			DCG += float32(1 / math.Log2(float64(i+2)))
			cnt++
		}
	}
	if cnt > 0 {
		return DCG / iDCG
	} else {
		return 0
	}
}

func GetNGCG(ctx context.Context, pool *pgxpool.Pool, mf *MF) (float32, float32, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Can't acquire connection from pull")
	}
	defer conn.Release()

	var testUsers []*User
	err = pgxscan.Select(
		ctx,
		conn,
		&testUsers,
		`
		SELECT users.user_id AS user_id, users.vec::float4[] AS vec
		FROM users INNER JOIN (SELECT user_id FROM test GROUP BY user_id) AS t
		ON users.user_id = t.user_id`)
	if err != nil {
		return 0, 0, errors.Wrap(err, "Can't select user_id from test")
	}

	var sumTrainNDCG, sumTestNDCG float32
	for _, u := range testUsers {
		var predict []bool
		err = conn.QueryRow(
			ctx,
			`
			WITH cur AS (
				SELECT (item_id IN (SELECT item_id FROM train WHERE user_id = $1)) AS in_train
				FROM items
				ORDER BY vec <#> $2::float4[]::vector)
			SELECT ARRAY_AGG(in_train) FROM cur
			`,
			u.Id,
			u.Vec).Scan(&predict)

		if err != nil {
			return 0, 0, errors.Wrap(err, "Can't read query")
		}
		sumTrainNDCG += NDCG(predict, 100)

		err = conn.QueryRow(
			ctx,
			`
			WITH
				tr AS (SELECT item_id FROM train WHERE user_id = $1),
				items_without_train AS (
					SELECT items.item_id AS item_id, items.vec AS vec
					FROM (items LEFT OUTER JOIN tr ON items.item_id = tr.item_id)
					WHERE tr.item_id is NULL
				),
				cur AS (
					SELECT (item_id IN (SELECT item_id FROM test WHERE user_id = $1)) AS in_test
					FROM items_without_train
					ORDER BY vec <#> $2::float4[]::vector)
			SELECT ARRAY_AGG(in_test) FROM cur
			`,
			u.Id,
			u.Vec).Scan(&predict)
		if err != nil {
			return 0, 0, errors.Wrap(err, "Can't read query")
		}
		sumTestNDCG += NDCG(predict, 100)

	}
	sumTrainNDCG /= float32(len(testUsers))
	sumTestNDCG /= float32(len(testUsers))

	return sumTrainNDCG, sumTestNDCG, nil
}

func (mf *MF) getTop(ctx context.Context, conn *pgx.Conn, userId string, topK int, onTest bool) (User, []*Item, error) {
	var user User
	err := pgxscan.Get(ctx, conn, &user, "SELECT user_id, vec::float4[], age, gender, occupation FROM users WHERE user_id = $1", userId)
	if err != nil {
		return User{}, []*Item{}, errors.Wrap(err, "Can't get user by id")
	}

	var items []*Item
	if onTest {
		err = pgxscan.Select(ctx, conn, &items, "SELECT item_id, vec::float4[] FROM items WHERE item_id NOT IN (SELECT item_id FROM train WHERE user_id = $1) ORDER BY vec <#> $2::float4[]::vector LIMIT $3", user.Id, user.Vec, topK)
	} else {
		err = pgxscan.Select(ctx, conn, &items, "SELECT item_id, vec::float4[] FROM items ORDER BY vec <#> $1::float4[]::vector LIMIT $2", user.Vec, topK)
	}

	if err != nil {
		return User{}, []*Item{}, errors.Wrap(err, "Can't get top items")
	}

	return user, items, nil
}

func (mf *MF) saveToFileTops(ctx context.Context, conn *pgx.Conn, topK int, onTest bool) error {
	var name string
	if onTest {
		name = "test"
	} else {
		name = "train"
	}

	var users []*User
	err := pgxscan.Select(ctx, conn, &users, "SELECT user_id FROM test GROUP BY user_id")
	if err != nil {
		return errors.Wrap(err, "Can't select users from test")
	}

	csvFile, err := os.Create(fmt.Sprintf("gbm/%s.csv", name))
	if err != nil {
		errors.Wrap(err, "Can't create csv file")
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	for _, cu := range users {
		u, items, err := mf.getTop(ctx, conn, cu.Id, 100, onTest)
		if err != nil {
			return errors.Wrap(err, "Can't get top")
		}
		var interact Interaction
		for _, i := range items {
			var row []string
			row = append(
				row,
				u.Id,
				fmt.Sprintf("%f", src.Dot(u.Vec, i.Vec)),
				fmt.Sprint(u.Age),
				u.Gender,
				u.Occupation)

			err = pgxscan.Get(ctx, conn, &interact, fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND item_id = $2", name), u.Id, i.Id)
			if pgxscan.NotFound(err) {
				row = append(row, "0")
			} else if err != nil {
				return errors.Wrap(err, "Error when get interaction")
			} else {
				row = append(row, "1")
			}
			csvwriter.Write(row)
		}
	}
	csvwriter.Flush()
	csvFile.Close()

	cdFile, err := os.Create(fmt.Sprintf("gbm/%s.cd", name))
	if err != nil {
		return errors.Wrap(err, "Can't create cd file")
	}
	defer cdFile.Close()

	cdFile.WriteString("0\tGroupId\tgroup_id\n")
	cdFile.WriteString("1\tNum\tproduct\n")
	cdFile.WriteString("2\tNum\tage\n")
	cdFile.WriteString("3\tCateg\tgender\n")
	cdFile.WriteString("4\tCateg\toccupation\n")
	cdFile.WriteString("5\tLabel\ty\n")

	return nil
}
