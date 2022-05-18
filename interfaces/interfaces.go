package interfaces

import "context"

type Migration func(ctx context.Context) error
