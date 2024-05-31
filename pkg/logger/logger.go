package logger

import (
	"github.com/google/logger"
	"os"
)

const logPath = "../../logs/log.txt"

var logFile *os.File

func InitializeLogger() error {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	logger.Init("LoggerApplication", true, true, logFile)
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logger.Info("Closing log file")
		err := logFile.Close()
		if err != nil {
			logger.Errorf("Failed to close log file: %v", err)
		} else {
			logger.Info("Log file closed")
		}
	}
	logger.Close()
}
