package logger

import "go.uber.org/zap/zapcore"

func LevelBySeverity(severity string, expected bool) zapcore.Level {
	switch severity {
	case "S1":
		return zapcore.ErrorLevel

	case "S2":
		if expected {
			return zapcore.InfoLevel
		}
		return zapcore.WarnLevel

	case "S3":
		if expected {
			return zapcore.DebugLevel
		}
		return zapcore.InfoLevel

	default:
		return zapcore.InfoLevel
	}
}
