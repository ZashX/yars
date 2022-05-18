package dataset

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/yars-ai/yars/src"
)

func TestModel(t *testing.T) {
	mf := MF{
		rng:        src.NewRandomGenerator(2),
		topK:       100,
		nFactors:   20,
		nSamples:   20,
		lr:         0.02,
		reg:        0,
		initMean:   0,
		initStdDev: 1,
	}

	ctx := context.Background()
	config, err := pgxpool.ParseConfig("postgres://user:password@localhost:5432/db")
	if err != nil {
		t.Fatal(err)
	}
	config.MaxConns = 20
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Release()

	err = mf.InitTableMovielens100k(ctx, conn.Conn())
	if err != nil {
		t.Fatal(err)
	}

	vTrain, vTest, err := GetNGCG(ctx, pool, &mf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(vTrain, vTest)

	if err = FitTrain(ctx, pool, &mf); err != nil {
		t.Fatal(err)
	}
	vTrain, vTest, err = GetNGCG(ctx, pool, &mf)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(vTrain, vTest)

	err = mf.saveToFileTops(ctx, conn.Conn(), 100, false)
	if err != nil {
		t.Fatal(err)
	}

	err = mf.saveToFileTops(ctx, conn.Conn(), 100, true)
	if err != nil {
		t.Fatal(err)
	}
}
