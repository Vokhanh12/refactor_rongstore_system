package logger

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type Logger interface {
	Info(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{})
	Error(ctx context.Context, msg string, err *aerrs.AppError, extra map[string]interface{})
	Warn(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{})
	Access(ctx context.Context, msg string, access AccessLog)
}
