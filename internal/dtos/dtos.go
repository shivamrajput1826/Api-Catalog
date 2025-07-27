package dtos

type CreateEventRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

type UpdateEventRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

type CreatePropertyRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Description string `json:"description"`
}

type TrackingPlanPropertyRequest struct {
	Name        string `json:"name" validate:"required"`
	Type        string `json:"type" validate:"required"`
	Required    bool   `json:"required"`
	Description string `json:"description"`
}

type TrackingPlanEventRequest struct {
	Name                 string                        `json:"name" validate:"required"`
	Description          string                        `json:"description"`
	Properties           []TrackingPlanPropertyRequest `json:"properties"`
	AdditionalProperties bool                          `json:"additionalProperties"`
	Type                 string                        `json:"type" validate:"required"`
}

type CreateTrackingPlanRequest struct {
	Name        string                     `json:"name" validate:"required"`
	Description string                     `json:"description"`
	Events      []TrackingPlanEventRequest `json:"events" validate:"required"`
}

type UpdateTrackingPlanRequest struct {
	Name        string                     `json:"name" validate:"required"`
	Description string                     `json:"description"`
	Events      []TrackingPlanEventRequest `json:"events" validate:"required"`
}

type ErrorResponse struct {
	Error   string      `json:"error"`
	Details interface{} `json:"details,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
