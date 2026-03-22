package logger

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{logger: logger}
}

func (z *ZapLogger) Log(ctx context.Context, level string, msg string, entry LogEntry, extra map[string]interface{}) {
	fields := buildFields(entry, extra)

	switch level {
	case "error":
		z.logger.Error(msg, fields...)
	case "warn":
		z.logger.Warn(msg, fields...)
	case "info":
		z.logger.Info(msg, fields...)
	default:
		z.logger.Debug(msg, fields...)
	}
}

func (z *ZapLogger) LogError(ctx context.Context, msg string, err *aerrs.AppError, extra map[string]interface{}) {
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

func (z *ZapLogger) LogAccess(ctx context.Context, msg string, access AccessLog) {
	fields := buildAccessFields(access)
	z.logger.Info(msg, fields...)
}

func LogAudit(ctx context.Context, msg string, extra map[string]interface{}) {
	fields := buildFields(ctx, extra)
	AuditLogger.Info(msg, fields...)
}
