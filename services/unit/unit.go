package unit

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"gitlab.com/yars-ai/yars/services/entity"
)

type Service interface {
	Create(ctx context.Context, in io.Reader) (io.Reader, error)
	Upsert(ctx context.Context, in io.Reader) (io.Reader, error)
	GetUnitById(ctx context.Context, id string) (interface{}, error)
}

type service struct {
	name              string
	repo              Store
	entityConstructor entity.EntityConstructor
}

func NewUnit(name string, r Store, c entity.EntityConstructor) *service {
	return &service{
		name:              name,
		repo:              r,
		entityConstructor: c,
	}
}

func (u *service) Create(ctx context.Context, in io.Reader) (io.Reader, error) {

	data := u.entityConstructor.NewSliceOfEntities()

	if err := json.NewDecoder(in).Decode(data); err != nil {
		return nil, fmt.Errorf("can't unmarshal: %w", err)
	}

	if err := u.repo.Create(ctx, data); err != nil {
		return nil, fmt.Errorf("can't create unit %s: %w", u.name, err)
	}

	return bytes.NewReader([]byte{}), nil
}

func (u *service) Upsert(ctx context.Context, in io.Reader) (io.Reader, error) {

	data := u.entityConstructor.NewSliceOfEntities()

	if err := json.NewDecoder(in).Decode(data); err != nil {
		return nil, fmt.Errorf("can't unmarshal: %w", err)
	}

	if err := u.repo.Upsert(ctx, data); err != nil {
		return nil, fmt.Errorf("can't upsert unit %s: %w", u.name, err)
	}

	return bytes.NewReader([]byte{}), nil
}

func (u *service) GetUnitById(ctx context.Context, id string) (interface{}, error) {
	unit := u.entityConstructor.New()
	if err := u.repo.GetUnitById(ctx, unit, id); err != nil {
		return nil, err
	}
	return unit, nil
}
