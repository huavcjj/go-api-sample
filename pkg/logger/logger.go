package logger

import (
	"go.uber.org/zap"
)

var (
	ZapLogger        *zap.Logger
	zapSugaredLogger *zap.SugaredLogger
)

func init() {
	cfg := zap.NewProductionConfig()
	ZapLogger = zap.Must(cfg.Build())
	zapSugaredLogger = ZapLogger.Sugar()
}

func Sync() {
	if err := zapSugaredLogger.Sync(); err != nil {
		ZapLogger.Error("Failed to sync logs", zap.Error(err))
	}
}

func Info(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Infow(msg, keysAndValues...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Debugw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Warnw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Errorw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Fatalw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	zapSugaredLogger.Panicw(msg, keysAndValues...)
}
