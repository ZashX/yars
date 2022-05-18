package dataset

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestLoadMovielens100k(t *testing.T) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/db")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	err = LoadMovielens100k(ctx, conn)
	if err != nil {
		t.Fatal(err)
	}
	err = DataTrainTestSplit(ctx, conn, 200, 1)
	if err != nil {
		t.Fatal(err)
	}
}
