package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func MustInitialize() {
	viper.SetConfigName("app.config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err.Error())
	}
}

func OnChange(cb func()) {
	viper.OnConfigChange(func(_ fsnotify.Event) {
		cb()
	})
}

func StartWatch() {
	viper.WatchConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
}
