package main

import (
	"app/config"
	"app/logger"
	"app/migrations"
	"app/ui"
	"log"
)

func main() {
	logger.Initialize()
	defer logger.Close()

	config.Initialize()

	err := migrations.Run(config.GetString("database.path"))

	if err != nil {
		log.Fatalf("Failed to executed migrations %s", err.Error())
	}

	var app = ui.NewApplication()

	config.OnChange(app.HandleConfigChange)
	config.StartWatch()

	app.Start()
}
