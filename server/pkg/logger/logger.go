package logger

import (
	"context"

	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	serviceInfo  ServiceInfo
	AccessLogger *zap.Logger
	ErrorLogger  *zap.Logger
	WarnLogger   *zap.Logger
	InfoLogger   *zap.Logger
	AuditLogger  *zap.Logger
)

type ServiceInfo struct {
	Name string
}

func GetServiceInfo() ServiceInfo {
	return serviceInfo
}

type RequestContext struct {
	TraceId  string
	UserId   string
	ClientId string
	RealmId  string
	SpanId   string
}

func (p RequestContext) IsEmpty() bool {
	return p.TraceId == "" &&
		p.UserId == "" &&
		p.ClientId == "" &&
		p.RealmId == ""
}

type LogEntry struct {
	ServiceInfo    ServiceInfo
	RequestContext RequestContext
	Code           string `json:"code"`
	Key            string `json:"key"`
	Expected       bool
	HTTPStatus     string `json:"http_status"`
	GRPCCode       string `json:"grpc_code"`
	Message        string `json:"message"`
	Cause          string `json:"cause"`
	Err            error
	ClientAction   string `json:"client_action"`
	ServerAction   string `json:"server_action"`
}

func Init(opts ...Option) error {

	for _, opt := range opts {
		opt(&serviceInfo)
	}

	encoderCfg := zapcore.EncoderConfig{TimeKey: "ts", LevelKey: "level", NameKey: "logger", CallerKey: "caller", MessageKey: "msg", StacktraceKey: "stack", LineEnding: zapcore.DefaultLineEnding, EncodeLevel: zapcore.LowercaseLevelEncoder, EncodeTime: zapcore.ISO8601TimeEncoder, EncodeDuration: zapcore.SecondsDurationEncoder, EncodeCaller: zapcore.ShortCallerEncoder}
	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	var err error

	AccessLogger, err = newFileLogger("logs/access.log", encoder, zap.InfoLevel)
	if err != nil {
		return err
	}

	InfoLogger, err = newFileLogger("logs/info.log", encoder, zap.InfoLevel)
	if err != nil {
		return err
	}

	ErrorLogger, err = newFileLogger("logs/error.log", encoder, zap.ErrorLevel)
	if err != nil {
		return err
	}

	WarnLogger, err = newFileLogger("logs/warn.log", encoder, zap.WarnLevel)
	if err != nil {
		return err
	}

	AuditLogger, err = newFileLogger("logs/audit.log", encoder, zap.InfoLevel)
	if err != nil {
		return err
	}

	return nil
}

func LogBySeverity(ctx context.Context, msg string, err *aerr.AppError) {

	if err == nil {
		LogWarn(ctx, msg, WarnParams{
			LogEntry: LogEntry{
				Message: "nil AppError passed to logger",
			},
		})
		return
	}

	level := LevelBySeverity(err.Severity, err.Expected)
	switch level {
	case zapcore.ErrorLevel:
		LogError(ctx, msg, ErrorParams{
			LogEntry: LogEntry{
				Code:     err.Code,
				Key:      err.Key,
				GRPCCode: err.GRPCCode,
				Message:  err.Message,
				Cause:    err.Cause,
			},
		})
	case zapcore.WarnLevel:
		LogWarn(ctx, msg, WarnParams{
			LogEntry: LogEntry{
				Code:     err.Code,
				Key:      err.Key,
				GRPCCode: err.GRPCCode,
				Message:  err.Message,
				Cause:    err.Cause,
				Err:      err.Err,
			},
		})
	case zapcore.InfoLevel:
		LogInfo(ctx, msg, InfoParams{
			LogEntry: LogEntry{
				Code:     err.Code,
				Key:      err.Key,
				GRPCCode: err.GRPCCode,
				Message:  err.Message,
				Cause:    err.Cause,
				Err:      err.Err,
			},
		})
	}
}
