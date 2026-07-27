package dispatcher

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func Wrap[T any, R any](
	fn func(
		context.Context,
		T,
	) (R, *aerrs.AppError),
) HandlerFunc {

	return func(
		ctx context.Context,
		payload any,
	) (any, *aerrs.AppError) {

		cmd, ok := payload.(T)
		if !ok {
			return nil, aerrs.New(core.INVALID_COMMAND_TYPE)
		}

		return fn(ctx, cmd)
	}
}
