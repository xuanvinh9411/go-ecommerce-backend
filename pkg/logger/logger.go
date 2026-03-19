package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	log = newLogger()
}


// func init() {
// 	instance,_ := zap.NewProduction()
// 	log = instance
// }


func newLogger() *zap.Logger {

	// Create a new logger instance
	logger, _ := zap.NewProduction()
	logger.Info("Logger initialized successfully")
	return logger	
}

func GetLogger() *zap.Logger {
	return log		
}