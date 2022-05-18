package ast

import (
	"bytes"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type FieldType interface {
	Node
	fieldTypeNode()
}

type Project struct {
	Name       Identifier
	Statements []Statement
}

func (p *Project) String() string {
	var out bytes.Buffer
	out.WriteString("project ")
	out.WriteString(p.Name.Value)
	out.WriteString(";")
	out.WriteString(" ")
	for i, stmt := range p.Statements {
		out.WriteString(stmt.String())
		if i != len(p.Statements)-1 {
			out.WriteString(" ")
		}
	}
	return out.String()
}

type UnitStatement struct {
	Name              Identifier
	FieldsDeclaration *FieldsDeclaration
}

func (us *UnitStatement) statementNode() {}
func (us *UnitStatement) String() string {
	var out bytes.Buffer
	out.WriteString("unit ")
	out.WriteString(us.Name.String())
	if us.FieldsDeclaration != nil {
		out.WriteString(us.FieldsDeclaration.String())
	}
	return out.String()
}

type ActionStatement struct {
	Name   Identifier
	Actor  Identifier
	Object Identifier
}

func (as *ActionStatement) statementNode() {}
func (as *ActionStatement) String() string {
	var out bytes.Buffer
	out.WriteString("action ")
	out.WriteString(as.Name.String())
	out.WriteString(" ")
	out.WriteString(as.Actor.String())
	out.WriteString("@")
	out.WriteString(as.Object.String())
	out.WriteString(";")
	return out.String()
}

type FieldsDeclaration struct {
	Fields []*Field
}

func (fd *FieldsDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("{ ")
	for i, s := range fd.Fields {
		out.WriteString(s.String())
		if i != len(fd.Fields)-1 {
			out.WriteString(" ")
		}
	}
	out.WriteString(" }")
	return out.String()
}

type Field struct {
	Name Identifier
	Type FieldType
}

func (f *Field) String() string {
	var out bytes.Buffer
	out.WriteString(f.Name.String())
	out.WriteString(" ")
	if f.Type != nil {
		out.WriteString(f.Type.String())
	}
	out.WriteString(";")
	return out.String()
}

type BaseType struct {
	Type Type
}

func (b *BaseType) fieldTypeNode() {}
func (b *BaseType) String() string {
	return string(b.Type)
}

type SliceType struct {
	BaseType *BaseType
}

func (a *SliceType) fieldTypeNode() {}
func (a *SliceType) String() string {
	return "[]" + a.BaseType.String()
}

type Identifier struct {
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string {
	return i.Value
}

type RecommendStatement struct {
	Name              Identifier
	Actor             Identifier
	Object            Identifier
	ParamsDeclaration *ParamsDeclaration
}

func (rec *RecommendStatement) statementNode() {}
func (rec *RecommendStatement) String() string {
	var out bytes.Buffer
	out.WriteString("recommend ")
	out.WriteString(rec.Name.Value)
	out.WriteString(" ")
	out.WriteString(rec.Actor.Value)
	out.WriteString("@")
	out.WriteString(rec.Object.Value)
	out.WriteString(" ")
	out.WriteString(rec.ParamsDeclaration.String())
	return out.String()
}

type Param struct {
	Name       Identifier
	Expression Expression
}

func (p *Param) String() string {
	var out bytes.Buffer
	out.WriteString(p.Name.Value)
	out.WriteString(": ")
	out.WriteString(p.Expression.String())
	return out.String()
}

type ParamsDeclaration struct {
	Params []*Param
}

func (d *ParamsDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("{ ")
	for _, p := range d.Params {
		out.WriteString(p.String())
		out.WriteString("; ")
	}
	out.WriteString("}")
	return out.String()
}

type Lvalue struct {
	Name Identifier
}

func (l *Lvalue) expressionNode() {}
func (l *Lvalue) String() string {
	var out bytes.Buffer
	out.WriteString(l.Name.Value)
	return out.String()
}

type LvalueField struct {
	Name  Identifier
	Field Identifier
}

func (l *LvalueField) expressionNode() {}
func (l *LvalueField) String() string {
	var out bytes.Buffer
	out.WriteString(l.Name.Value)
	out.WriteString(".")
	out.WriteString(l.Field.Value)
	return out.String()
}

type BinaryOp struct {
	Left     Expression
	Right    Expression
	Operator Operator
}

func (b *BinaryOp) expressionNode() {}
func (b *BinaryOp) String() string {
	var out bytes.Buffer
	out.WriteString(b.Left.String())
	out.WriteString(" ")
	out.WriteString(string(b.Operator))
	out.WriteString(" ")
	out.WriteString(b.Right.String())
	return out.String()
}

type UnaryOp struct {
	Expression Expression
	Operator   Operator
}

func (b *UnaryOp) expressionNode() {}
func (b *UnaryOp) String() string {
	var out bytes.Buffer
	out.WriteString(string(b.Operator))
	out.WriteString(" ")
	out.WriteString(b.Expression.String())
	return out.String()
}
