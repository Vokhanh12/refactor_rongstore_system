package usecases

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type Handler[T any] struct {
	Cond func(T) bool
	Exec func(context.Context, T) (any, *aerrs.AppError)
}

type MutateEngine[T any] struct {
	handlers []Handler[T]
}

func NewMutateEngine[T any](handlers []Handler[T]) *MutateEngine[T] {
	return &MutateEngine[T]{handlers: handlers}
}

func (e *MutateEngine[T]) Execute(ctx context.Context, items []Operation[T],
	buildResult func(opID string, data any, err *aerrs.AppError) dtos.MutateResultItem,
) []dtos.MutateResultItem {

	results := make([]dtos.MutateResultItem, 0, len(items))

	for _, item := range items {
		var (
			data any
			err  *aerrs.AppError
		)

		for _, h := range e.handlers {
			if h.Cond(item.Payload) {
				data, err = h.Exec(ctx, item.Payload)
				break
			}
		}

		if err == nil && data == nil {
			err = aerrs.New("MUTATE_OPERATION_UNSUPPORTED")
		}

		results = append(results, buildResult(item.OpID, data, err))
	}

	return results
}
