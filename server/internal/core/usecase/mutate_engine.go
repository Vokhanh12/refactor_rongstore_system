package usecases

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type MutateEngine[T any] struct {
	handlers []Handler[T]
}

func NewMutateEngine[T any](handlers []Handler[T]) *MutateEngine[T] {
	return &MutateEngine[T]{handlers: handlers}
}

func (e *MutateEngine[T]) Execute(ctx context.Context, items []Operation[T]) []dtos.MutateResultItemDTO {

	results := make([]dtos.MutateResultItemDTO, 0, len(items))

	for _, item := range items {

		// 1. Reject trước khi xử lý
		if !item.Success {
			results = append(results, mappers.BuildMutateResult(
				item.OpID,
				nil,
				aerrs.New(errs.MUTATE_OPERATION_REJECT),
			))
			continue
		}

		var (
			data    any
			err     *aerrs.AppError
			matched bool
		)

		// 2. Handler execution
		for _, h := range e.handlers {
			if h.Cond(item.Payload) {
				matched = true
				data, err = h.Exec(ctx, item.Payload)
				break
			}
		}

		// 3. No handler match
		if !matched {
			err = aerrs.New(errs.MUTATE_OPERATION_UNSUPPORTED)
		}

		results = append(results, mappers.BuildMutateResult(item.OpID, data, err))
	}

	return results
}
