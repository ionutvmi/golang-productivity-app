package main

import (
	"app/logger"
	"app/ui"
)

func main() {
	defer logger.Close()

	var app = ui.NewApplication()

	app.Start()
}
