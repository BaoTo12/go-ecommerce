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

	outputLog := lumberjack.Logger{
		Filename:   lg.LogFileName,
		MaxSize:    lg.MaxSize, // megabytes
		MaxBackups: lg.MaxBackup,
		MaxAge:     lg.MaxAge,   //days
		Compress:   lg.Compress, // disabled by default
	}

	encoder := getEncoderLog()
	sync := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stderr),
		zapcore.AddSync(&outputLog),
	)
	core := zapcore.NewCore(encoder, sync, level)

	return &LoggerZap{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}

}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encodeConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encodeConfig.EncodeName = zapcore.FullNameEncoder
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "Time"
	return zapcore.NewJSONEncoder(encodeConfig)
}
