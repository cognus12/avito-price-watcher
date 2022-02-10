package app

import (
	"apricescrapper/internal/avito"
	"errors"
	"log"
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
	log.Println("Init router")
	router := httprouter.New()

	avitoService := avito.NewService()
	handler := avito.NewHandler(avitoService)

	handler.Register(router)

	start(router)
}

func start(router http.Handler) {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("App started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			log.Println("server shutdown")
		default:
			log.Println(err)
		}
	}
}
