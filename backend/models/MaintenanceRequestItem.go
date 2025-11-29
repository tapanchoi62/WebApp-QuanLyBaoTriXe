package models

import "gorm.io/gorm"

type MaintenanceRequestItem struct {
	gorm.Model           `swaggerignore:"true"`
	MaintenanceRequestID uint    `json:"maintenanceRequestId"`
	ItemName             string  `json:"itemName"`
	Cost                 float64 `json:"cost"`
	Status               string  `json:"status"` // pending, done, cancelled
	Note                 string  `json:"note"`
}
