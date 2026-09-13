package logger

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
	Sync()
}

type noopLogger struct{}

func (n *noopLogger) Info(msg string, fields ...Field)  {}
func (n *noopLogger) Error(msg string, fields ...Field) {}
func (n *noopLogger) Fatal(msg string, fields ...Field) {}
func (n *noopLogger) Debug(msg string, fields ...Field) {}
func (n *noopLogger) Warn(msg string, fields ...Field)  {}
func (n *noopLogger) Sync()                             {}

var N Logger = &noopLogger{}

func NewLogger(out io.Writer) Logger {
	encCfg := zap.NewProductionEncoderConfig()
	encCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05")
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encCfg.CallerKey = ""
	encCfg.ConsoleSeparator = " "

	encoder := zapcore.NewConsoleEncoder(encCfg)
	core := zapcore.NewCore(encoder, zapcore.AddSync(out), zap.InfoLevel)

	return &logger{zap: zap.New(core)}
}

type logger struct {
	zap *zap.Logger
}

func (l *logger) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *logger) Sync() {
	_ = l.zap.Sync()
}

func String(key, value string) Field {
	return zap.String(key, value)
}

func Int(key string, value int) Field {
	return zap.Int(key, value)
}

func Float64(key string, value float64) Field {
	return zap.Float64(key, value)
}

func Float32(key string, value float32) Field {
	return zap.Float32(key, value)
}

func Bool(key string, value bool) Field {
	return zap.Bool(key, value)
}

func ErrorField(key string, err error) Field {
	return zap.NamedError(key, err)
}

func Reflect(key string, value any) Field {
	return zap.Any(key, value)
}
