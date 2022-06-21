package app

import (
	"apricescrapper/internal/advt"
	"apricescrapper/internal/config"
	"apricescrapper/internal/crawler"
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

	logger.Info("Connect to sqlite3, path: %v", cfg.DbPath)

	db, err := sqlite.New(schema, cfg.DbPath)

	if err != nil {
		logger.Panic(err.Error())
	}

	subscriptionRepository := subscription.NewRepository(db)

	subscriptionService := subscription.NewService(subscriptionRepository)
	advtService := advt.NewService(crawler)

	subscriptionHandler := subscription.NewHandler(subscriptionService, logger)
	advtHandler := advt.NewHandler(advtService, logger)

	subscriptionHandler.Register(router)
	advtHandler.Register(router)

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

	go shutdown.Gracefull(func() {
		server.Close()
		c.Close()
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
