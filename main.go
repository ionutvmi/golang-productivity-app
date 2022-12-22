package main

import (
	"app/config"
	"app/database"
	"app/logger"
	"app/migrations"
	"app/ui"
)

func main() {
	logger.MustInitialize()
	defer logger.Close()

	config.MustInitialize()

	var dbPath = config.GetString("database.path")

	migrations.MustRun(dbPath)

	database.MustInitialize(dbPath)
	defer database.Close()

	var app = ui.NewApplication()

	config.OnChange(app.HandleConfigChange)
	config.StartWatch()

	app.Start()
}
