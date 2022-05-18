package entity

import (
	"fmt"
	"reflect"

	"github.com/lib/pq"
	"gitlab.com/yars-ai/yars/ast"
)

type EntityConstructor interface {
	New() interface{}
	NewSliceOfEntities() interface{}
}

func NewEntityBuilder() *entityBuilder {
	return &entityBuilder{}
}

type entityBuilder struct {
	fields []*fieldConfig
}

type fieldConfig struct {
	name string
	typ  reflect.Type
	tag  string
}

type entityConstructor struct {
	definition reflect.Type
}

func (b *entityBuilder) AddField(name string, t ast.FieldType, tag string) error {
	typ, err := toReflectType(t)

	if err != nil {
		return err
	}

	// TODO sdelat adekvatno
	if t.String() == "[]float32" {
		tag += ` gorm:"type:float[]"`
	}

	if t.String() == "[]bool" {
		tag += ` gorm:"type:bool[]"`
	}

	if t.String() == "[]string" {
		tag += ` gorm:"type:text[]"`
	}

	b.fields = append(b.fields, &fieldConfig{
		name: name,
		typ:  typ,
		tag:  tag,
	})

	return nil
}

func (b *entityBuilder) BuildUnit() EntityConstructor {
	structFields := []reflect.StructField{
		{
			Name: "ID",
			Type: baseTypes[ast.String],
			Tag:  reflect.StructTag(`json:"ID" db:"id" gorm:"size:256;primaryKey"`),
		},
	}

	for _, field := range b.fields {
		structFields = append(structFields, reflect.StructField{
			Name: field.name,
			Type: field.typ,
			Tag:  reflect.StructTag(field.tag),
		})
	}

	return &entityConstructor{
		definition: reflect.StructOf(structFields),
	}
}

func (ec *entityConstructor) New() interface{} {
	return reflect.New(ec.definition).Interface()
}

func (ec *entityConstructor) NewSliceOfEntities() interface{} {
	return reflect.New(reflect.SliceOf(ec.definition)).Interface()
}

func toReflectType(t ast.FieldType) (reflect.Type, error) {
	switch ft := t.(type) {
	case *ast.BaseType:
		rt, ok := baseTypes[ft.Type]
		if !ok {
			return nil, fmt.Errorf("unsupported base type: %v", ft.Type)
		}
		return rt, nil
	case *ast.SliceType:
		if ft.BaseType == nil {
			return nil, fmt.Errorf("base type value for slice is nil")
		}
		rt, ok := sliceOfTypes[ft.BaseType.Type]
		if !ok {
			return nil, fmt.Errorf("unsupported base type: %v", ft.BaseType.Type)
		}
		return rt, nil
	default:
		return nil, fmt.Errorf("unsupportd field type: %v", ft)
	}
}

var baseTypes = map[ast.Type]reflect.Type{
	ast.Int32:   reflect.TypeOf(*new(int32)),
	ast.Float32: reflect.TypeOf(*new(float32)),
	ast.Bool:    reflect.TypeOf(*new(bool)),
	ast.String:  reflect.TypeOf(*new(string)),
}

var sliceOfTypes = map[ast.Type]reflect.Type{
	ast.Int32:   reflect.TypeOf(*new(pq.Int32Array)),
	ast.Float32: reflect.TypeOf(*new(pq.Float32Array)),
	ast.Bool:    reflect.TypeOf(*new(pq.BoolArray)),
	ast.String:  reflect.TypeOf(*new(pq.StringArray)),
}
