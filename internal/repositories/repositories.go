package repositories

import (
	"github.com/shivamrajput1826/api-catalog/internal/models"
	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (r *EventRepositoryImpl) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepositoryImpl) GetAll() ([]models.Event, error) {
	var events []models.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
func (r *EventRepositoryImpl) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepositoryImpl) Update(event *models.Event) error {
	return r.db.Save(event).Error
}

func (r *EventRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}

func (r *EventRepositoryImpl) GetByNameAndType(name, eventType string) (*models.Event, error) {
	var event models.Event
	if err := r.db.Where("name = ? AND type = ?", name, eventType).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

type PropertyRepositoryImpl struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) models.PropertyRepository {
	return &PropertyRepositoryImpl{db: db}
}

func (r *PropertyRepositoryImpl) Create(property *models.Property) error {
	return r.db.Create(property).Error
}

func (r *PropertyRepositoryImpl) GetAll() ([]models.Property, error) {
	var properties []models.Property
	err := r.db.Find(&properties).Error
	return properties, err
}

func (r *PropertyRepositoryImpl) GetByID(id uint) (*models.Property, error) {
	var property models.Property
	err := r.db.First(&property, id).Error
	if err != nil {
		return nil, err
	}
	return &property, nil
}

func (r *PropertyRepositoryImpl) Update(property *models.Property) error {
	return r.db.Save(property).Error
}

func (r *PropertyRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Property{}, id).Error
}

func (r *PropertyRepositoryImpl) GetByNameAndType(name, propertyType string) (*models.Property, error) {
	var property models.Property
	err := r.db.Where("name = ? AND type = ?", name, propertyType).First(&property).Error
	if err != nil {
		return nil, err
	}
	return &property, nil
}

type TrackingPlanRepositoryImpl struct {
	db *gorm.DB
}

func NewTrackingPlanRepository(db *gorm.DB) models.TrackingPlanRepository {
	return &TrackingPlanRepositoryImpl{db: db}
}

func (r *TrackingPlanRepositoryImpl) Create(plan *models.TrackingPlan) error {
	return r.db.Create(plan).Error
}

func (r *TrackingPlanRepositoryImpl) GetAll() ([]models.TrackingPlan, error) {
	var plans []models.TrackingPlan
	err := r.db.Preload("Events.Event").Preload("Events.Properties.Property").Find(&plans).Error
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (r *TrackingPlanRepositoryImpl) GetByID(id uint) (*models.TrackingPlan, error) {
	var plan models.TrackingPlan
	err := r.db.Preload("Events.Event").Preload("Events.Properties.Property").First(&plan, id).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func (r *TrackingPlanRepositoryImpl) Update(plan *models.TrackingPlan) error {
	return r.db.Save(plan).Error
}

func (r *TrackingPlanRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.TrackingPlan{}, id).Error
}

func (r *TrackingPlanRepositoryImpl) GetByName(name string) (*models.TrackingPlan, error) {
	var plan models.TrackingPlan
	err := r.db.Where("name = ?", name).First(&plan).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

type TransactionManagerImpl struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) models.TransactionManager {
	return &TransactionManagerImpl{db: db}
}

func (tm *TransactionManagerImpl) BeginTransaction() *gorm.DB {
	return tm.db.Begin()
}
