package logger

import (
	"fmt"
	"log"
)

type Logger interface {
	Info(msg string)
	Infof(format string, args ...any)
	Error(msg string)
	Errorf(format string, args ...any)
	Fatal(msg string)
	Fatalf(format string, args ...any)
}

const (
	InfoPrefix  = "INFO"
	ErrorPrefix = "ERROR"
	FatalPrefix = "FATAL"
)

func NewLogger() Logger {
	return &logger{}
}

type logger struct{}

func (l *logger) Info(msg string) {
	log.Printf("[%s] %s", InfoPrefix, msg)
}

func (l *logger) Infof(format string, args ...any) {
	log.Printf("[%s] %s", InfoPrefix, fmt.Sprintf(format, args...))
}

func (l *logger) Error(msg string) {
	log.Printf("[%s] %s", ErrorPrefix, msg)
}

func (l *logger) Errorf(format string, args ...any) {
	log.Printf("[%s] %s", ErrorPrefix, fmt.Sprintf(format, args...))
}

func (l *logger) Fatal(msg string) {
	log.Fatalf("[%s] %s", FatalPrefix, msg)
}

func (l *logger) Fatalf(format string, args ...any) {
	log.Fatalf("[%s] %s", FatalPrefix, fmt.Sprintf(format, args...))
}
