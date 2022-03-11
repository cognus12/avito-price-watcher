package shutdown

import (
	"apricescrapper/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

type Callback func()

func Gracefull(cb Callback) {

	logger := logger.GetInstance()

	signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, signals...)
	sig := <-sigc

	logger.Info("Caught signal %s. Shutting down...", sig)

	cb()
}
