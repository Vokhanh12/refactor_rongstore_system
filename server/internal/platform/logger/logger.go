package logger

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Logger interface {
	Log(ctx context.Context, level string, msg string, entry LogEntry, extra map[string]interface{})
	LogError(ctx context.Context, msg string, err *aerrs.AppError, extra map[string]interface{})
	LogAccess(ctx context.Context, msg string, access AccessLog)
}
