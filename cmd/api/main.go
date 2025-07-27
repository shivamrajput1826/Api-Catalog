// @title           API Catalog
// @version         1.0
// @description     This is the API Catalog service.
// @host            localhost:8080
// @BasePath        /
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/shivamrajput1826/api-catalog/cmd/api/docs"
	"github.com/shivamrajput1826/api-catalog/config"
	"github.com/shivamrajput1826/api-catalog/internal/db"
	"github.com/shivamrajput1826/api-catalog/internal/handlers"
	"github.com/shivamrajput1826/api-catalog/internal/routes"
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

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	if err := db.Migrate(database); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	defer db.Close(database)
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(middleware.RecoveryMiddleware)
	h := handlers.New(database)

	routes.Setup(app, h)
	PORT := config.GetConfigValue("PORT")

	error := app.Listen(":" + PORT)
	if error != nil {
		customLogger.Error("Error starting server", "error", error)
		panic(error)
	}

}
