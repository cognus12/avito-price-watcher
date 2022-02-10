package main

import (
	"apricescrapper/internal/app"
)

func main() {
	app := app.New()

	app.Run()
}

// TODO add logger
// TODO add config
// TODO add AppError middleware
// TODO improve parser
// TODO inject scrapper as dep to service
