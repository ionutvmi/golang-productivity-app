package main

import (
	"app/logger"
	"app/ui"
	"log"

	"github.com/spf13/viper"
)

func main() {
	defer logger.Close()

	viper.SetConfigName("app.config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}
	viper.WatchConfig()

	var app = ui.NewApplication()
	app.Start()
}
