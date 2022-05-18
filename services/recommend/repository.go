package recommend

import (
	"context"
)

type ActorStore interface {
}

type ObjectStore interface {
}

type UnitStoreReader interface {
	GetUnitById(ctx context.Context, id string) (interface{}, error)
}

type SquashActionStore interface {
	Upsert(ctx context.Context, actions []Action) error
	ScanAll(ctx context.Context) ([]Action, error)
}
