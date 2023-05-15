package logger

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

var (
	gLogger *zap.Logger
	err     error
	once    sync.Once
)

func Init() {
	once.Do(func() {
		cfg := zap.NewDevelopmentConfig()
		cfg.DisableStacktrace = true
		gLogger, err = cfg.Build(zap.AddCallerSkip(1))
		if err != nil {
			log.Println(err)
		}
	})
}

func D(msg string, fields ...zap.Field) {
	Init()
	gLogger.Debug(msg, fields...)
}

func I(msg string, fields ...zap.Field) {
	Init()
	gLogger.Info(msg, fields...)
}

func W(msg string, fields ...zap.Field) {
	Init()
	gLogger.Warn(msg, fields...)
}

func E(msg string, fields ...zap.Field) {
	Init()
	gLogger.Error(msg, fields...)
}

func F(msg string, fields ...zap.Field) {
	Init()
	gLogger.Fatal(msg, fields...)
}
