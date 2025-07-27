package db

import (
	"fmt"
	"time"

	"github.com/shivamrajput1826/api-catalog/config"
	"github.com/shivamrajput1826/api-catalog/internal/models"
	"github.com/shivamrajput1826/api-catalog/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var customLogger = logger.CreateLogger("database")

func ConnectDB() (*gorm.DB, error) {
	host := config.GetConfigValue("DATABASE.host")
	port := config.GetConfigValue("DATABASE.port")
	user := config.GetConfigValue("DATABASE.user")
	password := config.GetConfigValue("DATABASE.password")
	dbname := config.GetConfigValue("DATABASE.name")

	customLogger.Info(fmt.Sprintf("Attempting to connect to database: host=%s, port=%s, user=%s, dbname=%s",
		host, port, user, dbname))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		customLogger.Error("Failed to connect to database", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		customLogger.Error("Failed to get underlying sql.DB", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		customLogger.Error("Failed to ping database", err)
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	customLogger.Info("Successfully connected to database")
	return db, nil
}

func Migrate(db *gorm.DB) error {
	customLogger.Info("Running database migrations...")

	models := models.GetAllModels()

	customLogger.Info(fmt.Sprintf("Migrating %d models", len(models)))

	if err := db.AutoMigrate(models...); err != nil {
		customLogger.Error("Failed while migration", err)
		return err
	}

	customLogger.Info("Database migrations completed successfully")
	return nil
}

func Close(db *gorm.DB) error {
	customLogger.Info("Closing database connection...")
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
