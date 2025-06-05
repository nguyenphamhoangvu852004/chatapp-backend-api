package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugger := zap.NewExample()
	// sugger.Info("test")

	// dev, _ := zap.NewDevelopment()
	// dev.Debug("test")

	// proc, _ := zap.NewProduction()
	// proc.Info("test")

	encoderLog := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoderLog, sync, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("test", zap.String("test", "hello world"))
}

func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
