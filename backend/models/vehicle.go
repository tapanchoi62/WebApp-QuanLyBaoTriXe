package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID                  uint   `gorm:"primaryKey"`
	PlateNumber         string `gorm:"unique;not null"`
	Model               string
	Year                int
	Note                string
	MaintenanceRecords  []MaintenanceRecord
	MaintenanceRequests []MaintenanceRequest
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
