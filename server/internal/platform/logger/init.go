package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(serviceName string) (*ZapLogger, error) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(zapcore.Lock(zapcore.AddSync(&zapcore.BufferedWriteSyncer{}))),
		zapcore.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return NewZapLogger(logger), nil
}
