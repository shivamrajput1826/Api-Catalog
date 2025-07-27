package config

import (
	"os"

	"github.com/shivamrajput1826/api-catalog/logger"
	"github.com/spf13/viper"
)

var customLogger = logger.CreateLogger("config")

func LoadConfig() error {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	viper.SetConfigName(env)
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		customLogger.Error("Error reading config file", "error", err)
		return err
	}
	customLogger.Info("Configuration loaded successfully", "environment", env)
	return nil

}

func GetConfigValue(key string) string {
	return viper.GetString(key)
}
