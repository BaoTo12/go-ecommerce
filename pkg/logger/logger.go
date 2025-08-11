package logger

import (
	"os"

	"github.com/BaoTo12/go-ecommerce/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(lg *setting.LoggerSetting) *LoggerZap {
	var level zapcore.Level
	switch lg.LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	fileRotator := lumberjack.Logger{
		Filename:   lg.LogFileName,
		MaxSize:    lg.MaxSize, // megabytes
		MaxBackups: lg.MaxBackup,
		MaxAge:     lg.MaxAge,   //days
		Compress:   lg.Compress, // disabled by default
	}
	// --- encoders (console colored, file json)
	consoleEncoder, fileEncoder := getEncoderLog()

	// --- writers / syncers
	consoleWriter := zapcore.Lock(os.Stdout)
	fileWriter := zapcore.AddSync(&fileRotator)

	// --- cores
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, level)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, level)

	// combine so each log entry goes to both destinations
	core := zapcore.NewTee(consoleCore, fileCore)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	return &LoggerZap{logger}

}

// getEncoderLog returns two encoders:
//   - consoleEncoder: human-readable, colored output (for terminal)
//   - fileEncoder: JSON encoder (for files)
func getEncoderLog() (zapcore.Encoder, zapcore.Encoder) {
	// Console encoder: development-style with color for levels
	consoleEncCfg := zap.NewDevelopmentEncoderConfig()
	consoleEncCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // colorized levels
	consoleEncCfg.EncodeCaller = zapcore.ShortCallerEncoder
	consoleEncCfg.TimeKey = "Time"
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncCfg)

	// File encoder: production-style JSON (no ANSI color codes)
	fileEncCfg := zap.NewProductionEncoderConfig()
	fileEncCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncCfg.EncodeCaller = zapcore.ShortCallerEncoder
	fileEncCfg.TimeKey = "Time"
	fileEncoder := zapcore.NewJSONEncoder(fileEncCfg)

	return consoleEncoder, fileEncoder
}
