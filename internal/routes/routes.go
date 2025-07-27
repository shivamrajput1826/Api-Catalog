package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shivamrajput1826/api-catalog/internal/handlers"
)

func Setup(app *fiber.App, h *handlers.Handlers) {
	app.Get("/health", h.HealthCheck)

	api := app.Group("/api/v1")

	events := api.Group("/events")
	events.Post("/", h.CreateEvent)
	events.Get("/", h.GetEvents)
	events.Get("/:id", h.GetEvent)
	events.Put("/:id", h.UpdateEvent)
	events.Delete("/:id", h.DeleteEvent)

	properties := api.Group("/properties")
	properties.Post("/", h.CreateProperty)
	properties.Get("/", h.GetProperties)
	properties.Get("/:id", h.GetProperty)
	properties.Put("/:id", h.UpdateProperty)
	properties.Delete("/:id", h.DeleteProperty)

	trackingPlans := api.Group("/tracking-plans")
	trackingPlans.Post("/", h.CreateTrackingPlan)
	trackingPlans.Get("/", h.GetTrackingPlans)
	trackingPlans.Get("/:id", h.GetTrackingPlan)
	trackingPlans.Put("/:id", h.UpdateTrackingPlan)
	trackingPlans.Delete("/:id", h.DeleteTrackingPlan)

	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Route not found",
			"path":  c.Path(),
		})
	})
}
