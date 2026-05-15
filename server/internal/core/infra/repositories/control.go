package repositories

import "context"

type Locker[T any, ID any] interface {
	GetForUpdate(ctx context.Context, id ID) (*T, error)
}

type UnitOfWork interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}
