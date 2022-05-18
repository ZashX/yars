package action

import "context"

type Store interface {
	Upsert(ctx context.Context, data []Action) error
}
