package handlers

import (
	"github.com/shivamrajput1826/api-catalog/internal/dtos"
	"github.com/shivamrajput1826/api-catalog/internal/repositories"
	"github.com/shivamrajput1826/api-catalog/internal/services"
	"github.com/shivamrajput1826/api-catalog/internal/utils"
	"github.com/shivamrajput1826/api-catalog/internal/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handlers struct {
	eventService        *services.EventService
	propertyService     *services.PropertyService
	trackingPlanService *services.TrackingPlanService
}

func New(db *gorm.DB) *Handlers {
	eventRepo := repositories.NewEventRepository(db)
	propertyRepo := repositories.NewPropertyRepository(db)
	trackingPlanRepo := repositories.NewTrackingPlanRepository(db)
	txManager := repositories.NewTransactionManager(db)

	validator := validation.New()

	eventService := services.NewEventService(eventRepo, validator)
	propertyService := services.NewPropertyService(propertyRepo, validator)
	trackingPlanService := services.NewTrackingPlanService(trackingPlanRepo, eventRepo, propertyRepo, txManager, validator)

	return &Handlers{
		eventService:        eventService,
		propertyService:     propertyService,
		trackingPlanService: trackingPlanService,
	}
}

// Event Handlers
// CreateEvent godoc
// @Summary      Create a new event
// @Description  Create a new event with name, type, and description
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body  dtos.CreateEventRequest  true  "Event to create"
// @Success      201  {object}  models.Event
// @Failure      400  {object}  fiber.Map
// @Router       /events [post]
func (h *Handlers) CreateEvent(c *fiber.Ctx) error {
	var req dtos.CreateEventRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	event, err := h.eventService.CreateEvent(&req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

// GetEvents godoc
// @Summary      Get all events
// @Description  Retrieve a list of all events
// @Tags         events
// @Produce      json
// @Success      200  {array}  models.Event
// @Failure      500  {object}  fiber.Map
// @Router       /events [get]
func (h *Handlers) GetEvents(c *fiber.Ctx) error {
	events, err := h.eventService.GetAllEvents()
	if err != nil {
		return err
	}

	return c.JSON(events)
}

// GetEvent godoc
// @Summary      Get event by ID
// @Description  Retrieve a single event by its ID
// @Tags         events
// @Produce      json
// @Param        id   path      int  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /events/{id} [get]
func (h *Handlers) GetEvent(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	event, err := h.eventService.GetEventByID(id)
	if err != nil {
		return err
	}

	return c.JSON(event)
}

// UpdateEvent godoc
// @Summary      Update an event
// @Description  Update an event by its ID
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        id     path      int                      true  "Event ID"
// @Param        event  body      dtos.UpdateEventRequest  true  "Event update payload"
// @Success      200    {object}  models.Event
// @Failure      400    {object}  fiber.Map
// @Failure      404    {object}  fiber.Map
// @Router       /events/{id} [put]
func (h *Handlers) UpdateEvent(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	var req dtos.UpdateEventRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	event, err := h.eventService.UpdateEvent(id, &req)
	if err != nil {
		return err
	}

	return c.JSON(event)
}

// DeleteEvent godoc
// @Summary      Delete an event
// @Description  Delete an event by its ID
// @Tags         events
// @Param        id   path      int  true  "Event ID"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /events/{id} [delete]
func (h *Handlers) DeleteEvent(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.eventService.DeleteEvent(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CreateProperty godoc
// @Summary      Create a new property
// @Description  Create a new property with name, type, and description
// @Tags         properties
// @Accept       json
// @Produce      json
// @Param        property  body  dtos.CreatePropertyRequest  true  "Property to create"
// @Success      201  {object}  models.Property
// @Failure      400  {object}  fiber.Map
// @Router       /properties [post]
func (h *Handlers) CreateProperty(c *fiber.Ctx) error {
	var req dtos.CreatePropertyRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	property, err := h.propertyService.CreateProperty(&req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(property)
}

// GetProperties godoc
// @Summary      Get all properties
// @Description  Retrieve a list of all properties
// @Tags         properties
// @Produce      json
// @Success      200  {array}  models.Property
// @Failure      500  {object}  fiber.Map
// @Router       /properties [get]
func (h *Handlers) GetProperties(c *fiber.Ctx) error {
	properties, err := h.propertyService.GetAllProperties()
	if err != nil {
		return err
	}

	return c.JSON(properties)
}

// GetProperty godoc
// @Summary      Get property by ID
// @Description  Retrieve a single property by its ID
// @Tags         properties
// @Produce      json
// @Param        id   path      int  true  "Property ID"
// @Success      200  {object}  models.Property
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /properties/{id} [get]
func (h *Handlers) GetProperty(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	property, err := h.propertyService.GetPropertyByID(id)
	if err != nil {
		return err
	}

	return c.JSON(property)
}

// UpdateProperty godoc
// @Summary      Update a property
// @Description  Update a property by its ID
// @Tags         properties
// @Accept       json
// @Produce      json
// @Param        id        path      int                          true  "Property ID"
// @Param        property  body      dtos.UpdatePropertyRequest   true  "Property update payload"
// @Success      200       {object}  models.Property
// @Failure      400       {object}  fiber.Map
// @Failure      404       {object}  fiber.Map
// @Router       /properties/{id} [put]
func (h *Handlers) UpdateProperty(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	var req dtos.UpdatePropertyRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	property, err := h.propertyService.UpdateProperty(id, &req)
	if err != nil {
		return err
	}

	return c.JSON(property)
}

// DeleteProperty godoc
// @Summary      Delete a property
// @Description  Delete a property by its ID
// @Tags         properties
// @Param        id   path      int  true  "Property ID"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /properties/{id} [delete]
func (h *Handlers) DeleteProperty(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.propertyService.DeleteProperty(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// CreateTrackingPlan godoc
// @Summary      Create a new tracking plan
// @Description  Create a new tracking plan with events and properties
// @Tags         tracking-plans
// @Accept       json
// @Produce      json
// @Param        trackingPlan  body  dtos.CreateTrackingPlanRequest  true  "Tracking plan to create"
// @Success      201  {object}  models.TrackingPlan
// @Failure      400  {object}  fiber.Map
// @Router       /tracking-plans [post]
func (h *Handlers) CreateTrackingPlan(c *fiber.Ctx) error {
	var req dtos.CreateTrackingPlanRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	plan, err := h.trackingPlanService.CreateTrackingPlan(&req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(plan)
}

// GetTrackingPlans godoc
// @Summary      Get all tracking plans
// @Description  Retrieve a list of all tracking plans
// @Tags         tracking-plans
// @Produce      json
// @Success      200  {array}  models.TrackingPlan
// @Failure      500  {object}  fiber.Map
// @Router       /tracking-plans [get]
func (h *Handlers) GetTrackingPlans(c *fiber.Ctx) error {
	plans, err := h.trackingPlanService.GetAllTrackingPlans()
	if err != nil {
		return err
	}

	return c.JSON(plans)
}

// GetTrackingPlan godoc
// @Summary      Get tracking plan by ID
// @Description  Retrieve a single tracking plan by its ID
// @Tags         tracking-plans
// @Produce      json
// @Param        id   path      int  true  "Tracking Plan ID"
// @Success      200  {object}  models.TrackingPlan
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /tracking-plans/{id} [get]
func (h *Handlers) GetTrackingPlan(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	plan, err := h.trackingPlanService.GetTrackingPlanByID(id)
	if err != nil {
		return err
	}

	return c.JSON(plan)
}

// UpdateTrackingPlan godoc
// @Summary      Update a tracking plan
// @Description  Update a tracking plan by its ID
// @Tags         tracking-plans
// @Accept       json
// @Produce      json
// @Param        id            path      int                             true  "Tracking Plan ID"
// @Param        trackingPlan  body      dtos.UpdateTrackingPlanRequest  true  "Tracking plan update payload"
// @Success      200           {object}  models.TrackingPlan
// @Failure      400           {object}  fiber.Map
// @Failure      404           {object}  fiber.Map
// @Router       /tracking-plans/{id} [put]
func (h *Handlers) UpdateTrackingPlan(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	var req dtos.UpdateTrackingPlanRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON payload")
	}

	plan, err := h.trackingPlanService.UpdateTrackingPlan(id, &req)
	if err != nil {
		return err
	}

	return c.JSON(plan)
}

// DeleteTrackingPlan godoc
// @Summary      Delete a tracking plan
// @Description  Delete a tracking plan by its ID
// @Tags         tracking-plans
// @Param        id   path      int  true  "Tracking Plan ID"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  fiber.Map
// @Failure      404  {object}  fiber.Map
// @Router       /tracking-plans/{id} [delete]
func (h *Handlers) DeleteTrackingPlan(c *fiber.Ctx) error {
	id, err := utils.ParseUintID(c.Params("id"))
	if err != nil {
		return err
	}

	if err := h.trackingPlanService.DeleteTrackingPlan(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Returns the health status of the service
// @Tags         health
// @Produce      json
// @Success      200  {object}  fiber.Map
// @Router       /health [get]
func (h *Handlers) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "healthy",
		"service": "datacatalog",
		"version": "1.0.0",
	})
}
