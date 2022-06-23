package app

import (
	"apricescrapper/internal/advt"
	"apricescrapper/internal/config"
	"apricescrapper/internal/crawler"
	"apricescrapper/internal/observer"
	"apricescrapper/internal/subscription"
	"apricescrapper/pkg/logger"
	"apricescrapper/pkg/shutdown"
	"apricescrapper/pkg/sqlite"
	"errors"
	"fmt"
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
	logger := logger.GetInstance()
	cfg := config.GetConfig(logger)

	logger.Info("Init router")

	router := httprouter.New()

	crawler := crawler.Instance()

	logger.Info("Connecting to sqlite3, path: %v", cfg.DbPath)

	db, err := sqlite.New(schema, cfg.DbPath)

	if err != nil {
		logger.Panic(err.Error())
	}

	logger.Info("Successfully connected to sqlite3, path: %v", cfg.DbPath)

	subscriptionRepository := subscription.NewRepository(db)
	subscriptionService := subscription.NewService(subscriptionRepository)
	subscriptionHandler := subscription.NewHandler(subscriptionService, logger)

	advtService := advt.NewService(crawler)
	advtHandler := advt.NewHandler(advtService, logger)

	subscriptionHandler.Register(router)
	advtHandler.Register(router)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))

	if err != nil {
		logger.Fatal(err.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("App started on %s:%s", cfg.Host, cfg.Port)

	obs := observer.NewObserver(subscriptionService, advtService, logger)

	obs.Prepare().Run()

	go shutdown.Gracefull(func() {
		server.Close()
		crawler.Close()
	})

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Error("server shutdown")
		default:
			logger.Errorf(err)
		}
	}
}
