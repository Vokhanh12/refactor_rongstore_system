package usecases

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type Operation[T any] struct {
	OpID    string
	Payload T
	Error   *aerrs.AppError
}

type Handler[T any] struct {
	Cond func(T) bool
	Exec func(context.Context, T) (any, *aerrs.AppError)
}
