package usecases

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Operation[T any] struct {
	OpID    string
	Payload T
	Success bool
}

type Handler[T any] struct {
	Cond func(T) bool
	Exec func(context.Context, T) (any, *aerrs.AppError)
}
