package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var GlobalLogger *zap.Logger

func InitLogger(level string, development bool) error {
	loggerLevel := zap.InfoLevel

	switch level {
	case "debug":
		loggerLevel = zap.DebugLevel
	case "info":
		loggerLevel = zap.InfoLevel
	case "warn":
		loggerLevel = zap.WarnLevel
	case "error":
		loggerLevel = zap.ErrorLevel
	case "fatal":
		loggerLevel = zap.FatalLevel
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(loggerLevel),
		Development:      development,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	GlobalLogger, err = config.Build()
	if err != nil {
		return fmt.Errorf("config.Build: %w", err)
	}

	return nil
}
