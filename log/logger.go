package log

import (
	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Error("failed to sync logger", zap.Error(err))
			return
		}
	}(logger)
	return logger
}
