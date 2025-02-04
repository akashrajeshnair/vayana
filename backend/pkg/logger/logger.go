package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() *Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "timestamp"
	config.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02 15:04:05"))
	}

	consoleEncoder := zapcore.NewJSONEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Info(msg string, fields ...zapcore.Field) {
	l.Logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapcore.Field) {
	l.Logger.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zapcore.Field) {
	l.Logger.Fatal(msg, fields...)
}

func (l *Logger) With(fields ...zapcore.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
	}
}
