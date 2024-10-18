package logger

import (
	"github.com/difmaj/ms-credit-score/internal/pkg/config"
	"go.uber.org/zap"
)

// Logger is the logger instance
var Logger *zap.Logger

// Init initializes the logger.
func Init() {
	switch config.Env.Environment {
	case "production":
		Logger = zap.Must(zap.NewProduction())
	default:
		Logger = zap.Must(zap.NewDevelopment())
	}
}

// Sync flushes the logger.
func Sync() error {
	return Logger.Sync()
}
