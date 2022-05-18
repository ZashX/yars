package dataset

import (
	"context"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func LoadMovielens100k(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, "DROP TABLE IF EXISTS data")
	if err != nil {
		return errors.Wrap(err, "Can't drop table")
	}
	_, err = conn.Exec(ctx,
		`CREATE TABLE data (
			user_id varchar(256) NOT NULL,
			item_id varchar(256) NOT NULL,
			PRIMARY KEY(user_id, item_id)
		);`)
	if err != nil {
		return errors.Wrap(err, "Can't create table")
	}

	dataFile, err := os.Open("./movielens100k/u.data")
	if err != nil {
		return errors.Wrap(err, "Can't open file")
	}
	defer dataFile.Close()

	res, err := ioutil.ReadAll(dataFile)
	if err != nil {
		return errors.Wrap(err, "Can't read file")
	}

	rows := [][]interface{}{}
	dataLines := strings.Split(string(res), "\n")
	for _, l := range dataLines {
		vals := strings.Split(l, "\t")

		user_id := vals[0]
		item_id := vals[1]
		if val, _ := strconv.Atoi(vals[2]); val >= 4 {
			rows = append(rows, []interface{}{user_id, item_id})
		}
	}

	_, err = conn.CopyFrom(
		ctx,
		pgx.Identifier{"data"},
		[]string{"user_id", "item_id"},
		pgx.CopyFromRows(rows))
	if err != nil {
		return errors.Wrap(err, "Can't upload to pg data")
	}

	return nil
}

func DataTrainTestSplit(ctx context.Context, conn *pgx.Conn, numTestUsers int, numTestItems int) error {
	_, err := conn.Exec(ctx, "DROP TABLE IF EXISTS train")
	if err != nil {
		return errors.Wrap(err, "Can't drop table")
	}
	_, err = conn.Exec(ctx,
		`CREATE TABLE train (
			user_id varchar(256) NOT NULL,
			item_id varchar(256) NOT NULL,
			PRIMARY KEY(user_id, item_id)
		);`)
	if err != nil {
		return errors.Wrap(err, "Can't create table")
	}

	_, err = conn.Exec(ctx, "CREATE INDEX item_index ON train (item_id);")
	if err != nil {
		return errors.Wrap(err, "Can't create index for train")
	}

	_, err = conn.Exec(ctx, "DROP TABLE IF EXISTS test")
	if err != nil {
		return errors.Wrap(err, "Can't drop table")
	}
	_, err = conn.Exec(ctx,
		`CREATE TABLE test (
			user_id varchar(256) NOT NULL,
			item_id varchar(256) NOT NULL,
			PRIMARY KEY(user_id, item_id)
		);`)
	if err != nil {
		return errors.Wrap(err, "Can't create table")
	}

	rows, err := conn.Query(
		ctx,
		`
			WITH users AS (SElECT user_id FROM data GROUP BY user_id)
			SELECT user_id FROM users ORDER BY random() LIMIT $1
		`,
		numTestUsers)
	if err != nil {
		return err
	}
	defer rows.Close()

	testUserIds := []string{}
	for rows.Next() {
		var userId string
		err = rows.Scan(&userId)
		if err != nil {
			return err
		}
		testUserIds = append(testUserIds, userId)
		if err != nil {
			return err
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}

	for _, userId := range testUserIds {
		_, err = conn.Exec(
			ctx,
			`
				WITH cur AS (
					SELECT user_id, item_id
					FROM data
					WHERE data.user_id=$1
					ORDER BY random()
					LIMIT $2
				)
				INSERT INTO test SELECT * FROM cur		
			`,
			userId,
			numTestItems)
		if err != nil {
			return err
		}
	}

	_, err = conn.Query(
		ctx,
		`
			INSERT INTO train
			SELECT data.user_id, data.item_id
			FROM data
			LEFT OUTER JOIN test
			ON (data.user_id = test.user_id) AND (data.item_id = test.item_id)
			WHERE test.user_id is NULL
		`)
	if err != nil {
		return err
	}
	return nil
}
