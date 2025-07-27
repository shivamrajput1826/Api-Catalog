package db

import (
	"fmt"

	"github.com/shivamrajput1826/api-catalog/config"
	"github.com/shivamrajput1826/api-catalog/internal/models"
	"github.com/shivamrajput1826/api-catalog/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var customLogger = logger.CreateLogger("database")

func ConnectDB() (*gorm.DB, error) {
	host := config.GetConfigValue("DATABASE.host")
	port := config.GetConfigValue("DATABASE.port")
	user := config.GetConfigValue("DATABASE.user")
	password := config.GetConfigValue("DATABASE.password")
	dbname := config.GetConfigValue("DATABASE.name")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		customLogger.Error("Error reading config file", err)
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	customLogger.Info("Running database migrations...")

	models := models.GetAllModels()
	if err := db.AutoMigrate(models...); err != nil {
		customLogger.Error("Failed while migration", err)
		return err
	}

	customLogger.Info("Database migrations completed successfully")
	return nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
