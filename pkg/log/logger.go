package log

import (
	"github.com/workfoxes/gobase/pkg/config"
	"go.uber.org/zap"
)

var L *zap.Logger

func New(config *config.Config) *zap.Logger {
	if config.Env == "Dev" {
		L, _ = zap.NewDevelopment()
	} else {
		L, _ = zap.NewProduction()
	}
	return L
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	L.Debug(msg)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	L.Info(msg)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	L.Warn(msg)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	L.Error(msg)
}
