package config

import (
	"github.com/shivamrajput1826/api-catalog/logger"
	"github.com/spf13/viper"
)

var customLogger = logger.CreateLogger("config")

func LoadConfig() {
	viper.SetConfigName("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		customLogger.Error("Error reading config file", "error", err)
		panic(err)
	}
}

func GetConfigValue(key string) string {
	return viper.GetString(key)
}
