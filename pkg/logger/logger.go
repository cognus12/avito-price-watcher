package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)
}

type logger struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

func New() *logger {
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	return &logger{errorLogger, infoLogger}
}

func (l *logger) Info(msg string, v ...interface{}) {
	l.infoLogger.Printf(msg, v...)
}

func (l *logger) Error(msg string, v ...interface{}) {
	l.errorLogger.Printf(msg, v...)
}

func (l *logger) Fatal(msg string, v ...interface{}) {
	formatted := fmt.Sprintf(msg, v...)
	l.errorLogger.Fatal(formatted)
}

func (l *logger) Panic(msg string, v ...interface{}) {
	formatted := fmt.Sprintf(msg, v...)
	l.errorLogger.Panic(formatted)
}
