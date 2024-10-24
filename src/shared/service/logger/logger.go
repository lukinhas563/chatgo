package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

const (
	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL  = "LOG_LEVEL"
)

func init() {
	loggerConfig := zap.Config{
		OutputPaths: []string{getOutputLogs()},
		Level:       zap.NewAtomicLevelAt(getLevelLogs()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	logger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
}

func Info(message string, tags ...zap.Field) {
	logger.Info(message, tags...)
	logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	logger.Info(message, tags...)
	logger.Sync()
}

func getOutputLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))

	if output == "" {
		return "stdout"
	}

	return output
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
