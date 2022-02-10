package app

import (
	"apricescrapper/internal/avito"
	"apricescrapper/pkg/logger"
	"errors"
	"net"
	"net/http"
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

	avitoService := avito.NewService()
	handler := avito.NewHandler(avitoService, logger)

	handler.Register(router)

	start(router, logger)
}

func start(router http.Handler, logger logger.Logger) {
	listener, err := net.Listen("tcp", "localhost:3000")

	if err != nil {
		logger.Fatal(err.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("App started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Error("server shutdown")
		default:
			logger.Error(err.Error())
		}
	}
}
