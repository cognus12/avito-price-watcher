package app

import (
	"apricescrapper/internal/avito"
	"apricescrapper/internal/config"
	"apricescrapper/internal/crawler"
	"apricescrapper/pkg/logger"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
)

type App interface {
	Run()
}

type app struct{}

func New() *app {
	return &app{}
}

func (a *app) Run() {
	logger := logger.New()

	logger.Info("Init router")
	router := httprouter.New()

	crawler := crawler.NewCrawler()

	avitoService := avito.NewService(crawler)
	handler := avito.NewHandler(avitoService, logger)

	handler.Register(router)

	cfg := config.GetConfig(logger)

	start(router, logger, crawler, cfg)
}

func start(router http.Handler, logger logger.Logger, c crawler.Crawler, config *config.Config) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, config.Port))

	if err != nil {
		logger.Fatal(err.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("App started on %s:%s", config.Host, config.Port)

	go func() {
		signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}

		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, signals...)
		sig := <-sigc

		logger.Info("Caught signal %s. Shutting down...", sig)

		c.Stop()
		server.Close()
	}()

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Error("server shutdown")
		default:
			logger.Error(err.Error())
		}
	}

}

// func shutdown(s http.Server, c crawler.Crawler) {
// 	signals := []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}

// 	sigc := make(chan os.Signal, 1)
// 	signal.Notify(sigc, signals...)

// 	c.Stop()
// 	s.Close()
// }
