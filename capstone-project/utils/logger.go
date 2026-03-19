package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//var logger *zap.Logger

func InitLogger() (*zap.Logger, func(), error) {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	consoleWriter := zapcore.AddSync(os.Stdout)
	logFile, err := os.OpenFile("gin-RecipeApi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file: %w", err)
	}
	writer := zapcore.AddSync(logFile)
	defaultCore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, zap.DebugLevel),
		zapcore.NewCore(consoleEncoder, consoleWriter, zap.DebugLevel),
	)
	logger := zap.New(defaultCore, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	cleanup := func() {
		_ = logger.Sync()
		_ = logFile.Close()
	}
	return logger, cleanup, nil

}
