package test

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/pkg/errors"
// 	"gitlab.com/yars-ai/yars/ast"
// 	"gitlab.com/yars-ai/yars/repository"
// 	"gitlab.com/yars-ai/yars/services/action"
// 	"gitlab.com/yars-ai/yars/services/engine"
// 	"gitlab.com/yars-ai/yars/services/recommend"
// 	"gitlab.com/yars-ai/yars/services/recommend/boosting"
// 	"gitlab.com/yars-ai/yars/services/recommend/mf"
// 	"gitlab.com/yars-ai/yars/src"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func buildProject() (recommend.Service, error) {
// 	dsn := "host=localhost user=user password=password dbname=db port=5432"
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold:             time.Second,   // Slow SQL threshold
// 			LogLevel:                  logger.Silent, // Log level
// 			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
// 			Colorful:                  false,         // Disable color
// 		},
// 	)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		Logger:                 newLogger,
// 		SkipDefaultTransaction: true,
// 		PrepareStmt:            true,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrap(err, "can't open db connection")
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return nil, err
// 	}
// 	sqlDB.SetMaxOpenConns(5)

// 	project, err := ast.Parse(
// 		`
// 		project movielens;

// 		unit user {
// 			ID 			string;
// 			Age			int32;
// 			Gender 		string;
// 			Code		string;
// 			Occupation 	string;
// 		}

// 		unit item {
// 			ID 			string;
// 			Features    []float32;
// 		}
// 		`)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "can't parse project")
// 	}

// 	project.Statements = append(project.Statements, &ast.RecommendStatement{
// 		Name:   ast.Identifier{Value: "rec"},
// 		Actor:  ast.Identifier{Value: "user"},
// 		Object: ast.Identifier{Value: "item"},
// 	})

// 	engine, err := engine.NewBuilder(repository.NewFabric(db, "test"), project).Build()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't create engine")
// 	}
// 	actorUnit, err := engine.UnitByName("user")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't get user")
// 	}
// 	fileUser, err := os.Open("./movielens100k/user.json")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't read movielens100k/user.json")
// 	}
// 	_, err = actorUnit.Upsert(context.TODO(), io.Reader(fileUser))
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't create users")
// 	}
// 	objectUnit, err := engine.UnitByName("item")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't get item")
// 	}
// 	fileItem, err := os.Open("./movielens100k/item.json")
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't read movielens100k/item.json")
// 	}
// 	_, err = objectUnit.Upsert(context.TODO(), io.Reader(fileItem))
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Can't create items")
// 	}

// 	rec, _ := engine.RecommendByName("rec")
// 	return rec, nil
// }

// func TestTrainMovielens100k(t *testing.T) {
// 	rec, err := buildProject()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	retrainers := rec.GetRetrainers()
// 	var als *mf.ALS
// 	var cb *boosting.Catboost

// 	for _, r := range retrainers {
// 		switch s := r.(type) {
// 		case *mf.ALS:
// 			als = s
// 		case *boosting.Catboost:
// 			cb = s
// 		default:
// 			t.Fatal(fmt.Errorf("undefined type"))
// 		}
// 	}

// 	aFile, err := os.Open("movielens100k/action.json")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer aFile.Close()
// 	b, err := ioutil.ReadAll(aFile)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var actions []action.Action
// 	err = json.Unmarshal(b, &actions)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rng := src.NewRandomGenerator(1)
// 	src.RandomShuffle(&rng, actions)

// 	train := actions[:90000]
// 	test := actions[90000:]

// 	rmseTrain, err := recommend.RMSE(context.TODO(), als, train)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rmseTest, err := recommend.RMSE(context.TODO(), als, test)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	// fmt.Printf("RMSE on train %f, RMSE on test %f \n", rmseTrain, rmseTest)
// 	// err = rec.ApplyActions(context.TODO(), train)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rmseTrain, err = recommend.RMSE(context.TODO(), als, train)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// rmseTest, err = recommend.RMSE(context.TODO(), als, test)
// 	// if err != nil {
// 	// 	t.Fatal(err)
// 	// }
// 	// fmt.Printf("RMSE on train %f, RMSE on test %f \n", rmseTrain, rmseTest)

// 	err = cb.Retrain(context.TODO())
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rmseTrain, err = recommend.RMSE(context.TODO(), cb, train)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rmseTest, err = recommend.RMSE(context.TODO(), cb, test)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("RMSE on train %f, RMSE on test %f \n", rmseTrain, rmseTest)

// }
