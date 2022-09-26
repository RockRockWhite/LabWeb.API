package config

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("./config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
