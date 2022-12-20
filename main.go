package main

import (
	"app/config"
	"app/logger"
	"app/ui"
)

func main() {
	logger.Initialize()
	defer logger.Close()

	var app = ui.NewApplication()

	config.Initialize()
	config.OnChange(app.HandleConfigChange)
	config.StartWatch()

	app.Start()
}
