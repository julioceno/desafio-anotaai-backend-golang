package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func NewLogger() {
	Logger = _newLogger()
}

func NewLoggerWithPrefix(prefix string) *zap.Logger {
	logNew := _newLogger()
	if logNew == nil {
		return nil
	}

	return logNew.With(zap.String("prefix", prefix))
}

func _newLogger() *zap.Logger {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "message",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

func Info(message string, tags ...zap.Field) {
	Logger.Info(message, tags...)
	Logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	Logger.Error(message, tags...)
	Logger.Sync()
}

func Fatal(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	Logger.Fatal(message, tags...)
	Logger.Sync()
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
