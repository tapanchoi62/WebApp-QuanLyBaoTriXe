package models

import "gorm.io/gorm"

type MaintenanceLog struct {
	gorm.Model           `swaggerignore:"true"`
	MaintenanceRequestID uint    `json:"maintenanceRequestId"`
	Action               string  `json:"action"` // created, approved, in_progress, done, cancelled
	Description          string  `json:"description"`
	Cost                 float64 `json:"cost"`
	CreatedBy            string  `json:"createdBy"`
}
