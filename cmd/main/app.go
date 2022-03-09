package main

import (
	"apricescrapper/internal/app"
)

func main() {
	app := app.New()

	app.Run()
}

// TODO improve logger
// TODO shutdown func
