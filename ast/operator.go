package ast

type Operator string

const (
	And   = Type("and")
	Or    = "or"
	In    = "in"
	Plus  = "+"
	Minus = "-"
	Mul   = "*"
	Not   = "not"
)
