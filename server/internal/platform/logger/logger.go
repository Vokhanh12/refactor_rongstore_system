package logger

import (
	"context"

	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type Logger interface {
	Info(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{})
	Error(ctx context.Context, msg string, err dtos.InternalErrorDTO, extra map[string]interface{})
	Warn(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{})
	Access(ctx context.Context, msg string, access AccessLog)
}
