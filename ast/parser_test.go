package ast

import (
	"testing"
)

func TestUnit(t *testing.T) {
	tCase := `project yars; unit some{ str string; intField int32; BoolField bool; arr_str []string; } unit some2{ str string; intField int32; arr_float32 []float32; }`
	res, err := Parse(tCase)
	if err != nil {
		t.Fatal(err)
	}
	if res.String() != tCase {
		t.Fatalf("bad ast, expected:\n%s\ngot:\n%s", tCase, res.String())
	}
}

func TestAction(t *testing.T) {
	tCase := `project yars; action some a@b;`
	res, err := Parse(tCase)
	if err != nil {
		t.Fatal(err)
	}
	if res.String() != tCase {
		t.Fatalf("bad ast, expected:\n%s\ngot:\n%s", tCase, res.String())
	}
}

func TestRecommend(t *testing.T) {
	tCase := `project yars; recommend recname actor@object { rate: a * rate.b + c; filter: not (a and c); }`
	res, err := Parse(tCase)
	if err != nil {
		t.Fatal(err)
	}
	if res.String() != tCase {
		t.Fatalf("bad ast, expected:\n%s\ngot:\n%s", tCase, res.String())
	}
}
