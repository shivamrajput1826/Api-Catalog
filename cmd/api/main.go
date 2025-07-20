package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamrajput1826/api-catalog/config"
	"github.com/shivamrajput1826/api-catalog/logger"
	"github.com/shivamrajput1826/api-catalog/middleware"
)

var customLogger = logger.CreateLogger("API-Catalog")

func SetupRoutes(prefix string, app *fiber.App) {

}

func main() {
	config.LoadConfig()

	app := fiber.New(fiber.Config{
		BodyLimit:      1024 * 1024 * 10,
		Immutable:      true,
		ReadBufferSize: 1024 * 1024,
	})

	API_PREFIX := config.GetConfigValue("API_PREFIX")

	PORT := config.GetConfigValue("PORT")

	error := app.Listen(":" + PORT)
	if error != nil {
		customLogger.Error("Error starting server", "error", error)
		panic(error)
	}
	app.Use(middleware.RecoveryMiddleware)

}
