package mf

import (
	"context"
)

type ActorStore interface {
	CreateActorMF(ctx context.Context, actor Actor) error
	UpsertActorMF(ctx context.Context, actor Actor) error
	GetActorMFByID(ctx context.Context, id string) (Actor, error)
}

type ObjectStore interface {
	CreateObjectMF(ctx context.Context, object Object) error
	UpsertObjectMF(ctx context.Context, object Object) error
	GetObjectMFByID(ctx context.Context, id string) (Object, error)
	GetObjectsMF(ctx context.Context) ([]Object, error)
}
