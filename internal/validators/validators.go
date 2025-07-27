package validators

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shivamrajput1826/api-catalog/internal/dtos"
	"github.com/shivamrajput1826/api-catalog/logger"
)

var customLogger = logger.CreateLogger("validator")

var ValidEventTypes = map[string]bool{
	"track":    true,
	"identify": true,
	"alias":    true,
	"screen":   true,
	"page":     true,
}

var ValidPropertyTypes = map[string]bool{
	"string":  true,
	"number":  true,
	"boolean": true,
}

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateCreateEvent(req *dtos.CreateEventRequest) error {
	if req.Name == "" {
		customLogger.Error("ValidateCreateEventError", "name is required")
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Type == "" {
		customLogger.Error("ValidateCreateEventError", "type is required")
		return fiber.NewError(fiber.StatusBadRequest, "type is required")
	}
	if !ValidEventTypes[req.Type] {
		customLogger.Error("ValidateCreateEventError", "Wrong Validation type", req.Type)
		return fiber.NewError(fiber.StatusBadRequest,
			fmt.Sprintf("invalid event type '%s'. Must be one of: track, identify, alias, screen, page", req.Type))
	}
	return nil
}

func (v *Validator) ValidateUpdateEvent(req *dtos.UpdateEventRequest) error {
	if req.Name == "" {
		customLogger.Error("ValidateCreateEventError", "name is required")
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Type == "" {
		customLogger.Error("ValidateCreateEventError", "type is required")
		return fiber.NewError(fiber.StatusBadRequest, "type is required")
	}
	if !ValidEventTypes[req.Type] {
		customLogger.Error("ValidateCreateEventError", "Wrong Validation type", req.Type)
		return fiber.NewError(fiber.StatusBadRequest,
			fmt.Sprintf("invalid event type '%s'. Must be one of: track, identify, alias, screen, page", req.Type))
	}
	return nil
}

func (v *Validator) ValidateCreateProperty(req *dtos.CreatePropertyRequest) error {
	if req.Name == "" {
		customLogger.Error("ValidateCreatePropertyError", "name is required")
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Type == "" {
		customLogger.Error("ValidateCreatePropertyError", "type is required")
		return fiber.NewError(fiber.StatusBadRequest, "type is required")
	}
	if !ValidPropertyTypes[req.Type] {
		customLogger.Error("ValidateCreatePropertyError", "Wrong Validation type", req.Type)
		return fiber.NewError(fiber.StatusBadRequest,
			fmt.Sprintf("invalid property type '%s'. Must be one of: string, number, boolean", req.Type))
	}
	return nil
}

func (v *Validator) ValidateUpdateProperty(req *dtos.UpdatePropertyRequest) error {
	if req.Name == "" {
		customLogger.Error("ValidateUpdatePropertyError", "name is required")
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if req.Type == "" {
		customLogger.Error("ValidateUpdatePropertyError", "type is required")
		return fiber.NewError(fiber.StatusBadRequest, "type is required")
	}
	if !ValidPropertyTypes[req.Type] {
		customLogger.Error("ValidateUpdatePropertyError", "Wrong Validation type", req.Type)
		return fiber.NewError(fiber.StatusBadRequest,
			fmt.Sprintf("invalid property type '%s'. Must be one of: string, number, boolean", req.Type))
	}
	return nil
}

func (v *Validator) ValidateCreateTrackingPlan(req *dtos.CreateTrackingPlanRequest) error {
	if req.Name == "" {
		customLogger.Error("ValidateCreateTrackingPlanError", "name is required")
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if len(req.Events) == 0 {
		customLogger.Error("ValidateCreateTrackingPlanError", "events is required and cannot be empty")
		return fiber.NewError(fiber.StatusBadRequest, "events is required and cannot be empty")
	}

	for i, event := range req.Events {
		if event.Name == "" {
			customLogger.Error("ValidateCreateTrackingPlanError", fmt.Sprintf("event[%d].name is required", i))
			return fiber.NewError(fiber.StatusBadRequest,
				fmt.Sprintf("event[%d].name is required", i))
		}

		for j, prop := range event.Properties {
			if prop.Name == "" {
				customLogger.Error("ValidateCreateTrackingPlanError", fmt.Sprintf("event[%d].properties[%d].name is required", i, j))
				return fiber.NewError(fiber.StatusBadRequest,
					fmt.Sprintf("event[%d].properties[%d].name is required", i, j))
			}
			if prop.Type == "" {
				customLogger.Error("ValidateCreateTrackingPlanError", fmt.Sprintf("event[%d].properties[%d].type is required", i, j))
				return fiber.NewError(fiber.StatusBadRequest,
					fmt.Sprintf("event[%d].properties[%d].type is required", i, j))
			}
			if !ValidPropertyTypes[prop.Type] {
				customLogger.Error("ValidateCreateTrackingPlanError", fmt.Sprintf("event[%d].properties[%d].type '%s' is invalid", i, j, prop.Type))
				return fiber.NewError(fiber.StatusBadRequest,
					fmt.Sprintf("event[%d].properties[%d].type '%s' is invalid. Must be one of: string, number, boolean", i, j, prop.Type))
			}
		}
	}

	return nil
}

func (v *Validator) ValidateUpdateTrackingPlan(req *dtos.UpdateTrackingPlanRequest) error {
	return v.ValidateCreateTrackingPlan((*dtos.CreateTrackingPlanRequest)(req))
}

func (v *Validator) ValidateID(id string) error {
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id parameter is required")
	}
	return nil
}
