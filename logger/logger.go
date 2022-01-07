package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderconfig := zap.NewProductionEncoderConfig()
	encoderconfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = SyslogTimeEncoder
	config.EncoderConfig = encoderconfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan  2 15:04:05"))
}

func Info(message string, anyfield ...zap.Field) {
	log.Info(message, anyfield...)
}

func Debug(message string, anyfield ...zap.Field) {
	log.Debug(message, anyfield...)
}

func Error(message string, anyfield ...zap.Field) {
	log.Error(message, anyfield...)
}
