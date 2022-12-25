package main

import (
	"app/config"
	"app/database"
	"app/logger"
	"app/migrations"
	"app/ui"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "cpu-profile" {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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
