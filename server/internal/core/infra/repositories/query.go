package repositories

import "context"

type Exists[Q any] interface {
	Exists(ctx context.Context, query Q) (bool, error)
}

type Get[T any, Q any] interface {
	Get(ctx context.Context, query Q) (*T, error)
}

type List[T any, Q any] interface {
	List(ctx context.Context, query Q) ([]T, error)
}

type Count[Q any] interface {
	Count(ctx context.Context, query Q) (int64, error)
}
