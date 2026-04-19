package logger

import (
	"context"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	logger     *zap.Logger
	LogContext LogContext
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {

	logcontext := LogContext{
		ServiceName: "service_name",
		TraceID:     "trace_id",
		UserID:      "user_id",
		ClientID:    "client_id",
		RealmID:     "realm_id",
		SpanID:      "span_id",
	}

	return &ZapLogger{
		logger:     logger,
		LogContext: logcontext,
	}
}

// Access implements [Logger].
func (z *ZapLogger) Access(ctx context.Context, msg string, access AccessLog) {
	fields := buildAccessFields(access)
	z.logger.Info(msg, fields...)
}

// Error implements [Logger].
func (z *ZapLogger) Error(ctx context.Context, msg string, errDTO commonv1.InternalAppErrorDTO, extra map[string]interface{}) {

	entry := LogEntry{
		Context:      z.LogContext,
		Code:         errDTO.Code,
		Key:          errDTO.Key,
		Message:      errDTO.Message,
		Cause:        errDTO.Cause,
		CauseDetail:  errDTO.CauseDetail,
		ClientAction: errDTO.ClientAction,
		ServerAction: errDTO.ServerAction,
		Expected:     errDTO.Expected,
		GRPCCode:     errDTO.GRPCCode,
	}

	level := LevelBySeverity(errDTO.Severity, errDTO.Expected)
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

// Info implements [Logger].
func (z *ZapLogger) Info(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{}) {
	fields := buildFields(entry, extra)
	z.logger.Info(msg, fields...)
}

// Warn implements [Logger].
func (z *ZapLogger) Warn(ctx context.Context, msg string, entry LogEntry, extra map[string]interface{}) {
	fields := buildFields(entry, extra)
	z.logger.Warn(msg, fields...)
}
