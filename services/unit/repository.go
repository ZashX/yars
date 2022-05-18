package unit

import "context"

type Store interface {
	Create(ctx context.Context, data interface{}) error
	Upsert(ctx context.Context, data interface{}) error
	GetUnitById(ctx context.Context, unit interface{}, id string) error
}
