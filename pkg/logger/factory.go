package logger

import (
	"github.com/alphazento/betgo/pkg/container"
	"github.com/alphazento/betgo/pkg/logger/zap"
)

type Logger struct {
}

func Factory() interface{} {
	registerDrivers()
	l := Logger{}
	return &l
}

func registerDrivers() {
	container.Bind("zap_logger", zap.Factory, true)
	// container.Bind('logrus', Zap.Factory, true)
}

func (c *Logger) Info(message string) {
	logger := container.Get("zap_logger")
	if log, ok := logger.(ILogger); ok {
		log.Info(message)
	}
}

func (c *Logger) Infof(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Infof(message, args...)
	}
}

func (c *Logger) Errorf(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Errorf(message, args...)
	}
}

func (c *Logger) Fatal(message string) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Fatal(message)
	}
}

func (c *Logger) Fatalf(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Fatalf(message, args...)
	}
}

func (c *Logger) Warnf(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Warnf(message, args...)
	}
}

func (c *Logger) Debug(message string) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Debug(message)
	}
}

func (c *Logger) Debugf(message string, args ...interface{}) {
	if logger, ok := container.Get("zap_logger").(ILogger); ok {
		logger.Debugf(message, args...)
	}
}
