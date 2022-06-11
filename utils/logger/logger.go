package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	log, _ = zap.NewProduction()
	defer log.Sync()
	log.Info("logging service started")
}

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Error(msg, tags...)
	log.Sync()
}

func Panic(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.DPanic(msg, tags...)
	log.Sync()
}
