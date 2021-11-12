package log

import (
	"os"
	"sync"

	_logger "github.com/sirupsen/logrus"
)

var L _logger.Logger

type Logger struct {
	once     sync.Once
	_default *_logger.Logger
}

func Init() *_logger.Logger {
	return &_logger.Logger{
		Out:          os.Stderr,
		Formatter:    new(_logger.TextFormatter),
		Hooks:        make(_logger.LevelHooks),
		Level:        _logger.TraceLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	L.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	L.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	L.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	L.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	L.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	L.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	L.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	L.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	L.Fatal(args...)
}
