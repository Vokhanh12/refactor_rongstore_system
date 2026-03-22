package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(serviceName string) (*ZapLogger, error) {

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		MessageKey:     "msg",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	accessFile := &lumberjack.Logger{
		Filename:   "logs/access.log",
		MaxSize:    50, // MB
		MaxBackups: 7,
		MaxAge:     30, // days
		Compress:   true,
	}

	infoFile := &lumberjack.Logger{
		Filename:   "logs/info.log",
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	errorFile := &lumberjack.Logger{
		Filename:   "logs/error.log",
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	warnFile := &lumberjack.Logger{
		Filename:   "logs/warn.log",
		MaxSize:    50,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	accessCore := zapcore.NewCore(encoder, zapcore.AddSync(accessFile), zapcore.InfoLevel)
	infoCore := zapcore.NewCore(encoder, zapcore.AddSync(infoFile), zapcore.InfoLevel)
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorFile), zapcore.ErrorLevel)
	warnCore := zapcore.NewCore(encoder, zapcore.AddSync(warnFile), zapcore.WarnLevel)

	combined := zapcore.NewTee(accessCore, infoCore, errorCore, warnCore)
	logger := zap.New(combined, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return NewZapLogger(logger), nil
}
