package models

import (
	"gorm.io/gorm"
)

type Event struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;index:idx_event_name_type,unique"`
	Type        string `json:"type" gorm:"not null;index:idx_event_name_type,unique"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime  int64  `json:"update_time" gorm:"autoUpdateTime"`
}

type Property struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;index:idx_property_name_type,unique"`
	Type        string `json:"type" gorm:"not null;index:idx_property_name_type,unique"`
	Description string `json:"description"`
	CreateTime  int64  `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime  int64  `json:"update_time" gorm:"autoUpdateTime"`
}

type TrackingPlan struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	Name        string              `json:"name" gorm:"not null;unique"`
	Description string              `json:"description"`
	Events      []TrackingPlanEvent `json:"events" gorm:"foreignKey:TrackingPlanID;constraint:OnDelete:CASCADE"`
	CreateTime  int64               `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime  int64               `json:"update_time" gorm:"autoUpdateTime"`
}

type TrackingPlanEvent struct {
	ID                   uint                        `json:"id" gorm:"primaryKey"`
	TrackingPlanID       uint                        `json:"tracking_plan_id"`
	EventID              uint                        `json:"event_id"`
	Event                Event                       `json:"event" gorm:"foreignKey:EventID"`
	AdditionalProperties bool                        `json:"additionalProperties"`
	Properties           []TrackingPlanEventProperty `json:"properties" gorm:"foreignKey:TrackingPlanEventID;constraint:OnDelete:CASCADE"`
}

type TrackingPlanEventProperty struct {
	ID                  uint     `json:"id" gorm:"primaryKey"`
	TrackingPlanEventID uint     `json:"tracking_plan_event_id"`
	PropertyID          uint     `json:"property_id"`
	Property            Property `json:"property" gorm:"foreignKey:PropertyID"`
	Required            bool     `json:"required"`
}

func GetAllModels() []interface{} {
	return []interface{}{
		&Event{},
		&Property{},
		&TrackingPlan{},
		&TrackingPlanEvent{},
		&TrackingPlanEventProperty{},
	}
}

type EventRepository interface {
	Create(event *Event) error
	GetAll() ([]Event, error)
	GetByID(id uint) (*Event, error)
	Update(event *Event) error
	Delete(id uint) error
	GetByNameAndType(name, eventType string) (*Event, error)
}

type PropertyRepository interface {
	Create(property *Property) error
	GetAll() ([]Property, error)
	GetByID(id uint) (*Property, error)
	Update(property *Property) error
	Delete(id uint) error
	GetByNameAndType(name, propertyType string) (*Property, error)
}

type TrackingPlanRepository interface {
	Create(plan *TrackingPlan) error
	GetAll() ([]TrackingPlan, error)
	GetByID(id uint) (*TrackingPlan, error)
	Update(plan *TrackingPlan) error
	Delete(id uint) error
	GetByName(name string) (*TrackingPlan, error)
}

type TransactionManager interface {
	BeginTransaction() *gorm.DB
}
