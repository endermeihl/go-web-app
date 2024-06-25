package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	Config = viper.New()
	Config.SetConfigFile(".env")
	Config.AutomaticEnv()

	if err := Config.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}
