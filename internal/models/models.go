package model

import "time"

// event
// property
// tracking plan
// now tracking plan has a name and description..
// for each tracking plan.. definitely there are events..
// for each event.. it has some properties.. if i m not wrong..

type EventType string

const (
	EventTypeTrack    EventType = "track"
	EventTypeIdentify EventType = "identify"
	EventTypeAlias    EventType = "alias"
	EventTypeScreen   EventType = "screen"
	EventTypePage     EventType = "page"
)

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Type        EventType `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Properties  []Property `gorm:"many2many:event_properties;"`
}

type PropertyType string

const (
	PropertyTypeString  PropertyType = "string"
	PropertyTypeNumber  PropertyType = "number"
	PropertyTypeBoolean PropertyType = "boolean"
)

type Property struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"not null"`
	Type        PropertyType `gorm:"not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TrackingPlan struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Events      []TrackingPlanEvent
}

type TrackingPlanEvent struct {
	ID                   uint `gorm:"primaryKey"`
	TrackingPlanID       uint
	EventID              uint
	AdditionalProperties bool
	Properties           []TrackingPlanEventProperty
	Event                Event
}

type TrackingPlanEventProperty struct {
	ID                  uint `gorm:"primaryKey"`
	TrackingPlanEventID uint
	PropertyID          uint
	Required            bool
	Property            Property
}
