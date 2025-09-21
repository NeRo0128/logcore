package logcore

import "logcore/logger"

var defaultLogger = logger.NewLogger()

func LogSuccess(msg string) {
	defaultLogger.Info(msg)
}

func LogInfo(msg string) {
	defaultLogger.Info(msg)
}

func LogWarning(msg string) {
	defaultLogger.Warn(msg)
}

func LogError(msg string) {
	defaultLogger.Error(msg)
}
