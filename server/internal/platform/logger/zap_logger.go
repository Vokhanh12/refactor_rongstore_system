package logger

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: logger}
}

func (z *ZapLogger) Info(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{}) {
	fields := buildFields(entry, extra)
	z.logger.Info(msg, fields...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{}) {
	fields := buildFields(entry, extra)
	z.logger.Warn(msg, fields...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, err *aerrs.AppError, extra map[string]interface{}) {
	entry := FromAppError(err)
	level := LevelBySeverity(err.Severity, err.Expected)
	fields := buildFields(entry, extra)

	switch level {
	case zapcore.ErrorLevel:
		z.logger.Error(msg, fields...)
	case zapcore.WarnLevel:
		z.logger.Warn(msg, fields...)
	default:
		z.logger.Info(msg, fields...)
	}
}

func (z *ZapLogger) Access(ctx context.Context, msg string, access AccessLog) {
	fields := buildAccessFields(access)
	z.logger.Info(msg, fields...)
}
