package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := encodeConfig()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("This is an info message", zap.Int("Line", 1))
	logger.Error("This is an error message", zap.Int("Line", 2))
}

func encodeConfig() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	// Set the time format to ISO8601 1234567890 -> 2006-01-02T15:04:05Z07:00
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//covert key ts -> time hoặc bất cứ gì
	encodeConfig.TimeKey = "time"

	//from ino -> caller
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	//"caller" : "main.go:10" file ghi ở dòng nào
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encodeConfig)
}
func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
