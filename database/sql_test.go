package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v4"
)

func TestVector(t *testing.T) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/db")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, "CREATE EXTENSION IF NOT EXISTS vector;")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "DROP TABLE IF EXISTS test_table;")
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(
		ctx,
		`
		CREATE TABLE test_table (
			user_id varchar(256) NOT NULL,
			vec vector(3) NOT NULL,
			PRIMARY KEY(user_id)
		);`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "INSERT INTO test_table VALUES ($1, $2::float4[])", "a", []float32{1, 0.5, 0.5})
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "INSERT INTO test_table VALUES ($1, $2::float4[])", "b", []float32{0.5, 1, 0.5})
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec(ctx, "INSERT INTO test_table VALUES ($1, $2::float4[])", "c", []float32{0.5, 0.5, 1})
	if err != nil {
		t.Fatal(err)
	}

	var id string
	err = conn.QueryRow(ctx, "SELECT user_id FROM test_table ORDER BY vec <#> $1::float4[]::vector LIMIT 1", []float32{0.5, 1, 0}).Scan(&id)
	if err != nil {
		t.Fatal(err)
	}
	if id != "b" {
		t.Fatal(fmt.Sprint(id, " not equal b"))
	}

	_, err = conn.Exec(ctx, "DROP TABLE test_table;")
	if err != nil {
		t.Fatal(err)
	}
}
