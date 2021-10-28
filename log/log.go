package log

import (
	"fmt"

	"go.uber.org/zap"
	"team1.asia/fibo/config"
)

var Zap *zap.Logger

// Setup Zap logger.
func SetupLogger() {
	var (
		err    error
		zapCfg zap.Config
	)

	switch config.App.Env {
	case "production":
		zapCfg = zap.NewProductionConfig()
	default:
		zapCfg = zap.NewDevelopmentConfig()
	}

	zapCfg.OutputPaths = []string{
		"stdout",
		fmt.Sprintf(config.App.Log.FilePath, config.App.Server.Name, config.App.Env),
	}

	Zap, err = zapCfg.Build()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer Zap.Sync()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	Zap.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	Zap.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	Zap.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	Zap.Error(msg, fields...)
}
