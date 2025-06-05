package main

import (
	"os"

	"go.uber.org/zap/zapcore"
)

func main() {
	// sugger := zap.NewExample()
	// sugger.Info("test")

	// dev, _ := zap.NewDevelopment()
	// dev.Debug("test")

	// proc, _ := zap.NewProduction()
	// proc.Info("test")

}


func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
