package ast

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "gitlab.com/yars-ai/yars/antlr"
)

func Parse(code string) (*Project, error) {
	is := antlr.NewInputStream(code)
	lexer := parser.NewYarsLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewYarsParser(stream)
	l := yarsListener{}
	antlr.ParseTreeWalkerDefault.Walk(&l, p.Project())
	n, ok := l.pop()
	if !ok {
		return nil, fmt.Errorf("expect Porject ast")
	}
	project, ok := n.(*Project)
	if !ok {
		return nil, fmt.Errorf("can't parse code, expected *ast.Project, got %T", n)
	}
	return project, nil
}

type yarsListener struct {
	*parser.BaseYarsListener
	stack []Node
}

func (y *yarsListener) VisitErrorNode(node antlr.ErrorNode) {
	log.Fatalf("error node %v", node.GetText()) // todo заменить на panic
}

func (y *yarsListener) ExitProject(ctx *parser.ProjectContext) {
	p := &Project{
		Name: Identifier{Value: ctx.IDENTIFIER().GetText()},
	}
	for n, ok := y.pop(); ok; n, ok = y.pop() {
		if s, ok := n.(Statement); !ok {
			log.Fatalf("expected statement declaration, got %T", n) // todo заменить на panic
		} else {
			p.Statements = append([]Statement{s}, p.Statements...)
		}
	}
	y.push(p)
}

func (y *yarsListener) ExitUnitStatement(ctx *parser.UnitStatementContext) {
	n, ok := y.pop()
	if !ok {
		log.Fatalf("expected FieldsDeclaration") // todo заменить на panic
	}

	fields, ok := n.(*FieldsDeclaration)
	if !ok {
		log.Fatalf("expected fields declaration, got %T", n) // todo заменить на panic
	}

	u := &UnitStatement{
		Name:              Identifier{Value: ctx.IDENTIFIER().GetText()},
		FieldsDeclaration: fields,
	}

	y.push(u)
}

func (y *yarsListener) ExitActionStatement(ctx *parser.ActionStatementContext) {
	a := &ActionStatement{
		Name:   Identifier{Value: ctx.GetName().GetText()},
		Actor:  Identifier{Value: ctx.GetActor().GetText()},
		Object: Identifier{Value: ctx.GetObject().GetText()},
	}

	y.push(a)
}

func (y *yarsListener) ExitFieldsDecl(ctx *parser.FieldsDeclContext) {
	fd := &FieldsDeclaration{}
	for n, ok := y.pop(); ok; n, ok = y.pop() {
		if f, ok := n.(*Field); !ok {
			y.push(n)
			break
		} else {
			fd.Fields = append([]*Field{f}, fd.Fields...)
		}
	}

	y.push(fd)
}

func (y *yarsListener) ExitFieldDecl(ctx *parser.FieldDeclContext) {
	n, ok := y.pop()
	if !ok {
		log.Fatalf("expected type") // todo заменить на panic
	}

	ft, ok := n.(FieldType)
	if !ok {
		log.Fatalf("expected FieldType, got %T", n) // todo заменить на panic
	}

	y.push(&Field{
		Name: Identifier{Value: ctx.IDENTIFIER().GetText()},
		Type: ft,
	})
}

func (y *yarsListener) ExitArrayType(ctx *parser.ArrayTypeContext) {
	n, ok := y.pop()
	if !ok {
		log.Fatalf("expected base type") // todo заменить на panic
	}

	baseType, ok := n.(*BaseType)
	if !ok {
		log.Fatalf("unsupported array type") // todo заменить на panic
	}

	arr := SliceType{BaseType: baseType}
	y.push(&arr)
}

func (y *yarsListener) ExitBaseType(ctx *parser.BaseTypeContext) { //TODO исправить на типы на основе parsera
	switch ctx.GetText() {
	case "string":
		t := BaseType{Type: String}
		y.push(&t)
	case "int32":
		t := BaseType{Type: Int32}
		y.push(&t)
	case "float32":
		t := BaseType{Type: Float32}
		y.push(&t)
	case "bool":
		t := BaseType{Type: Bool}
		y.push(&t)
	default:
		log.Fatalf("undefined base type name: %s", ctx.GetText()) // todo исправить на добавление ошибки в парсер
	}
}

func (y *yarsListener) ExitRecommendStatement(c *parser.RecommendStatementContext) {
	n, ok := y.pop()
	if !ok {
		panic("expected ParamsDeclaration")
	}

	pd, ok := n.(*ParamsDeclaration)
	if !ok {
		panic(fmt.Sprintf("expected ParamsDeclaration, got %T", n))
	}

	rs := RecommendStatement{
		Name:              Identifier{Value: c.GetName().GetText()},
		Actor:             Identifier{Value: c.GetActor().GetText()},
		Object:            Identifier{Value: c.GetObject().GetText()},
		ParamsDeclaration: pd,
	}
	y.push(&rs)
}

func (y *yarsListener) ExitParamsDecl(c *parser.ParamsDeclContext) {
	pd := ParamsDeclaration{}
	for n, ok := y.pop(); ok; n, ok = y.pop() {
		if f, ok := n.(*Param); !ok {
			y.push(n)
			break
		} else {
			pd.Params = append([]*Param{f}, pd.Params...)
		}
	}
	y.push(&pd)
}

func (y *yarsListener) ExitParamDecl(c *parser.ParamDeclContext) {
	n, ok := y.pop()
	if !ok {
		panic("expected expression")
	}

	exp, ok := n.(Expression)
	if !ok {
		panic("unsupported expression type")
	}

	p := Param{
		Name:       Identifier{Value: c.GetName().GetText()},
		Expression: exp,
	}

	y.push(&p)
}

func (y *yarsListener) ExitLvalue(c *parser.LvalueContext) {
	name := c.GetName()
	field := c.GetField()
	if field == nil {
		y.push(&Lvalue{
			Name: Identifier{Value: name.GetText()},
		})
	} else {
		y.push(&LvalueField{
			Name:  Identifier{Value: name.GetText()},
			Field: Identifier{Value: field.GetText()},
		})
	}
}

func (y *yarsListener) ExitBinaryOp(c *parser.BinaryOpContext) {
	n, ok := y.pop()
	if !ok {
		panic("expected expression")
	}

	right, ok := n.(Expression)
	if !ok {
		panic("unsupported expression type")
	}

	n, ok = y.pop()
	if !ok {
		panic("expected expression")
	}

	left, ok := n.(Expression)
	if !ok {
		panic("unsupported expression type")
	}

	bo := BinaryOp{
		Left:     left,
		Right:    right,
		Operator: Operator(c.GetOp().GetText()),
	}

	y.push(&bo)
}

func (y *yarsListener) ExitUnaryOp(c *parser.UnaryOpContext) {
	n, ok := y.pop()
	if !ok {
		panic("expected expression")
	}

	exp, ok := n.(Expression)
	if !ok {
		panic("unsupported expression type")
	}

	uo := UnaryOp{
		Expression: exp,
		Operator:   Operator(c.GetOp().GetText()),
	}

	y.push(&uo)
}

func (y *yarsListener) push(n Node) {
	y.stack = append(y.stack, n)
}

func (y *yarsListener) pop() (Node, bool) {
	if len(y.stack) < 1 {
		return nil, false
	}
	result := y.stack[len(y.stack)-1]
	y.stack = y.stack[:len(y.stack)-1]

	return result, true
}
