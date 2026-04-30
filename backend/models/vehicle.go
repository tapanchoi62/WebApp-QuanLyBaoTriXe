package models

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID                  uint           `gorm:"primaryKey" json:"ID"`
	PlateNumber         string         `gorm:"unique;not null" json:"plate_number"`
	Model               string         `gorm:"column:model" json:"model_vehicle"`
	Year                int            `json:"year"`
	Note                string         `json:"note"`
	MaintenanceRecords  []MaintenanceRecord  `json:"-"`
	MaintenanceRequests []MaintenanceRequest `json:"-"`
	CreatedAt           time.Time      `json:"CreatedAt"`
	UpdatedAt           time.Time      `json:"UpdatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
}
