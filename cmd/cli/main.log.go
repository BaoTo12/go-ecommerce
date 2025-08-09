package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// ! 1 Log example with sugar and logger
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d", "Chi Bao", 24)

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Chi Bao"), zap.Int("age", 24))

	// ! 2 - Test different log level
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// ! 3 Test custom logger
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core)

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// formatting logs
func getEncoderLog() zapcore.Encoder {
	// Default production log: {"level":"info","ts":1754731769.2190158,"caller":"cli/main.log.go:19","msg":"Hello NewProduction"}
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1754731769.2190158 --> 2025-08-09T16:29:29.218+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// change key 'ts'
	encodeConfig.TimeKey = "Time Stamp"
	// change "info" to "INFO"
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:19"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// This function writes log to both console and file
func getWriterSync() zapcore.WriteSyncer {
	err := os.MkdirAll("./log", 0755)
	if err != nil {
		panic(err) // Handle this properly in production
	}
	// create or open a file
	// os.ModePerm allows full access
	file, err := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// AddSync: the function transforms both destinations into objects that Zap can work with.
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}

// How to do custom logger
// Components
/**
	1. Encoder determines whether logs appear as JSON, console-friendly text or other formats
	2. Writer syncer manages where your logs actually go
	3. The core combines these elements with a log level filter, creating your complete logging engine.

**/
