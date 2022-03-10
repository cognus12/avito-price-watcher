package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger interface {
	Info(msg string, v ...interface{})
	Error(msg string, v ...interface{})
	Fatal(msg string, v ...interface{})
	Panic(msg string, v ...interface{})
}

type logger struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
}

var lock = &sync.Mutex{}

var instance *logger

func initialize() {
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)

	instance = &logger{errorLogger, infoLogger}
}

func GetInstance() *logger {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			initialize()
			return instance
		}
	}

	return instance
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
