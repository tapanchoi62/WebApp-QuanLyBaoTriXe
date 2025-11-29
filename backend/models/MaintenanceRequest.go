package models

import (
	"time"

	"gorm.io/gorm"
)

type MaintenanceRequest struct {
	gorm.Model  `swaggerignore:"true"`
	VehicleID   uint      `json:"vehicleId"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // pending, in_progress, done
	RequestDate time.Time `json:"requestDate"`
}
