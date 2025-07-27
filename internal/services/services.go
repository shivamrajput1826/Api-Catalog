package services

import (
	"github.com/shivamrajput1826/api-catalog/internal/dtos"
	"github.com/shivamrajput1826/api-catalog/internal/models"
	"github.com/shivamrajput1826/api-catalog/internal/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EventService struct {
	eventRepo models.EventRepository
	validator *validation.Validator
}

func NewEventService(eventRepo models.EventRepository, validator *validation.Validator) *EventService {
	return &EventService{
		eventRepo: eventRepo,
		validator: validator,
	}
}

func (s *EventService) CreateEvent(req *dtos.CreateEventRequest) (*models.Event, error) {
	if err := s.validator.ValidateCreateEvent(req); err != nil {
		return nil, err
	}

	event := &models.Event{
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
	}

	if err := s.eventRepo.Create(event); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create event")
	}

	return event, nil
}

func (s *EventService) GetAllEvents() ([]models.Event, error) {
	events, err := s.eventRepo.GetAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch events")
	}
	return events, nil
}

func (s *EventService) GetEventByID(id uint) (*models.Event, error) {
	event, err := s.eventRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Event not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch event")
	}
	return event, nil
}

func (s *EventService) UpdateEvent(id uint, req *dtos.UpdateEventRequest) (*models.Event, error) {
	if err := s.validator.ValidateUpdateEvent(req); err != nil {
		return nil, err
	}

	event, err := s.eventRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Event not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch event")
	}

	event.Name = req.Name
	event.Type = req.Type
	event.Description = req.Description

	if err := s.eventRepo.Update(event); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update event")
	}

	return event, nil
}

func (s *EventService) DeleteEvent(id uint) error {
	if err := s.eventRepo.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete event")
	}
	return nil
}

type PropertyService struct {
	propertyRepo models.PropertyRepository
	validator    *validation.Validator
}

func NewPropertyService(propertyRepo models.PropertyRepository, validator *validation.Validator) *PropertyService {
	return &PropertyService{
		propertyRepo: propertyRepo,
		validator:    validator,
	}
}

func (s *PropertyService) CreateProperty(req *dtos.CreatePropertyRequest) (*models.Property, error) {
	if err := s.validator.ValidateCreateProperty(req); err != nil {
		return nil, err
	}

	property := &models.Property{
		Name:        req.Name,
		Type:        req.Type,
		Description: req.Description,
	}

	if err := s.propertyRepo.Create(property); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create property")
	}

	return property, nil
}

func (s *PropertyService) GetAllProperties() ([]models.Property, error) {
	properties, err := s.propertyRepo.GetAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch properties")
	}
	return properties, nil
}

func (s *PropertyService) GetPropertyByID(id uint) (*models.Property, error) {
	property, err := s.propertyRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Property not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch property")
	}
	return property, nil
}

func (s *PropertyService) UpdateProperty(id uint, req *dtos.UpdatePropertyRequest) (*models.Property, error) {
	if err := s.validator.ValidateUpdateProperty(req); err != nil {
		return nil, err
	}

	property, err := s.propertyRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Property not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch property")
	}

	property.Name = req.Name
	property.Type = req.Type
	property.Description = req.Description

	if err := s.propertyRepo.Update(property); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update property")
	}

	return property, nil
}

func (s *PropertyService) DeleteProperty(id uint) error {
	if err := s.propertyRepo.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete property")
	}
	return nil
}

type TrackingPlanService struct {
	trackingPlanRepo models.TrackingPlanRepository
	eventRepo        models.EventRepository
	propertyRepo     models.PropertyRepository
	txManager        models.TransactionManager
	validator        *validation.Validator
}

func NewTrackingPlanService(
	trackingPlanRepo models.TrackingPlanRepository,
	eventRepo models.EventRepository,
	propertyRepo models.PropertyRepository,
	txManager models.TransactionManager,
	validator *validation.Validator,
) *TrackingPlanService {
	return &TrackingPlanService{
		trackingPlanRepo: trackingPlanRepo,
		eventRepo:        eventRepo,
		propertyRepo:     propertyRepo,
		txManager:        txManager,
		validator:        validator,
	}
}

func (s *TrackingPlanService) CreateTrackingPlan(req *dtos.CreateTrackingPlanRequest) (*models.TrackingPlan, error) {
	if err := s.validator.ValidateCreateTrackingPlan(req); err != nil {
		return nil, err
	}

	tx := s.txManager.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	trackingPlan := &models.TrackingPlan{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := tx.Create(trackingPlan).Error; err != nil {
		tx.Rollback()
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create tracking plan")
	}

	for _, eventReq := range req.Events {
		event, err := s.findOrCreateEvent(tx, eventReq.Name, eventReq.Type, eventReq.Description)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		trackingPlanEvent := &models.TrackingPlanEvent{
			TrackingPlanID:       trackingPlan.ID,
			EventID:              event.ID,
			AdditionalProperties: eventReq.AdditionalProperties,
		}

		if err := tx.Create(trackingPlanEvent).Error; err != nil {
			tx.Rollback()
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create tracking plan event")
		}

		for _, propReq := range eventReq.Properties {
			property, err := s.findOrCreateProperty(tx, propReq.Name, propReq.Type, propReq.Description)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			trackingPlanEventProperty := &models.TrackingPlanEventProperty{
				TrackingPlanEventID: trackingPlanEvent.ID,
				PropertyID:          property.ID,
				Required:            propReq.Required,
			}

			if err := tx.Create(trackingPlanEventProperty).Error; err != nil {
				tx.Rollback()
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create tracking plan event property")
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to commit transaction")
	}

	result, err := s.trackingPlanRepo.GetByID(trackingPlan.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch created tracking plan")
	}

	return result, nil
}

func (s *TrackingPlanService) GetAllTrackingPlans() ([]models.TrackingPlan, error) {
	plans, err := s.trackingPlanRepo.GetAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch tracking plans")
	}
	return plans, nil
}

func (s *TrackingPlanService) GetTrackingPlanByID(id uint) (*models.TrackingPlan, error) {
	plan, err := s.trackingPlanRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Tracking plan not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch tracking plan")
	}
	return plan, nil
}

func (s *TrackingPlanService) UpdateTrackingPlan(id uint, req *dtos.UpdateTrackingPlanRequest) (*models.TrackingPlan, error) {
	if err := s.validator.ValidateUpdateTrackingPlan(req); err != nil {
		return nil, err
	}

	tx := s.txManager.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	trackingPlan, err := s.trackingPlanRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fiber.NewError(fiber.StatusNotFound, "Tracking plan not found")
		}
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch tracking plan")
	}

	trackingPlan.Name = req.Name
	trackingPlan.Description = req.Description

	if err := tx.Save(trackingPlan).Error; err != nil {
		tx.Rollback()
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to update tracking plan")
	}

	if err := tx.Where("tracking_plan_id = ?", trackingPlan.ID).Delete(&models.TrackingPlanEvent{}).Error; err != nil {
		tx.Rollback()
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete existing events")
	}

	for _, eventReq := range req.Events {
		event, err := s.findOrCreateEvent(tx, eventReq.Name, "track", eventReq.Description)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		trackingPlanEvent := &models.TrackingPlanEvent{
			TrackingPlanID:       trackingPlan.ID,
			EventID:              event.ID,
			AdditionalProperties: eventReq.AdditionalProperties,
		}

		if err := tx.Create(trackingPlanEvent).Error; err != nil {
			tx.Rollback()
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create tracking plan event")
		}

		for _, propReq := range eventReq.Properties {
			property, err := s.findOrCreateProperty(tx, propReq.Name, propReq.Type, propReq.Description)
			if err != nil {
				tx.Rollback()
				return nil, err
			}

			trackingPlanEventProperty := &models.TrackingPlanEventProperty{
				TrackingPlanEventID: trackingPlanEvent.ID,
				PropertyID:          property.ID,
				Required:            propReq.Required,
			}

			if err := tx.Create(trackingPlanEventProperty).Error; err != nil {
				tx.Rollback()
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create tracking plan event property")
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to commit transaction")
	}

	result, err := s.trackingPlanRepo.GetByID(trackingPlan.ID)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch updated tracking plan")
	}

	return result, nil
}

func (s *TrackingPlanService) DeleteTrackingPlan(id uint) error {
	if err := s.trackingPlanRepo.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete tracking plan")
	}
	return nil
}

func (s *TrackingPlanService) findOrCreateEvent(tx *gorm.DB, name, eventType, description string) (*models.Event, error) {
	var event models.Event
	if err := tx.Where("name = ? AND type = ?", name, eventType).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			event = models.Event{
				Name:        name,
				Type:        eventType,
				Description: description,
			}
			if err := tx.Create(&event).Error; err != nil {
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create event")
			}
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to query event")
		}
	} else {
		if event.Description != description && description != "" {
			return nil, fiber.NewError(fiber.StatusConflict, "Event exists with different description")
		}
	}
	return &event, nil
}

func (s *TrackingPlanService) findOrCreateProperty(tx *gorm.DB, name, propertyType, description string) (*models.Property, error) {
	if !validation.ValidPropertyTypes[propertyType] {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid property type. Must be one of: string, number, boolean")
	}

	var property models.Property
	if err := tx.Where("name = ? AND type = ?", name, propertyType).First(&property).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			property = models.Property{
				Name:        name,
				Type:        propertyType,
				Description: description,
			}
			if err := tx.Create(&property).Error; err != nil {
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create property")
			}
		} else {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to query property")
		}
	} else {
		if property.Description != description && description != "" {
			return nil, fiber.NewError(fiber.StatusConflict, "Property exists with different description")
		}
	}
	return &property, nil
}
