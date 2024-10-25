package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestLoggerConfig_DefaultOutput(t *testing.T) {
	os.Setenv(LOG_OUTPUT, "")
	defer os.Unsetenv(LOG_OUTPUT)

	output := getOutputLogs()
	assert.Equal(t, "stdout", output, "Expected default output to be stdout")
}

func TestLoggerConfig_CustomOutput(t *testing.T) {
	os.Setenv(LOG_OUTPUT, "stderr")
	defer os.Unsetenv(LOG_OUTPUT)

	output := getOutputLogs()
	assert.Equal(t, "stderr", output, "Expected output to be stderr")
}

func TestLoggerConfig_LogLevel_debug(t *testing.T) {
	os.Setenv(LOG_LEVEL, "debug")
	defer os.Unsetenv(LOG_LEVEL)

	level := getLevelLogs()
	assert.Equal(t, zapcore.DebugLevel, level, "Expected log level to be debug")
}

func TestLoggerConfig_LogLevel_error(t *testing.T) {
	os.Setenv(LOG_LEVEL, "error")
	defer os.Unsetenv(LOG_LEVEL)

	level := getLevelLogs()
	assert.Equal(t, zapcore.ErrorLevel, level, "Expected log level to be error")
}
