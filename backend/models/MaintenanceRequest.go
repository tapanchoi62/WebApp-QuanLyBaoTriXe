package models

import "time"

type MaintenanceRequest struct {
	ID         uint   `gorm:"primaryKey"`
	VehicleID  uint   `gorm:"not null"`
	Status     string `gorm:"default:'Pending'"` // Pending, Approved, Rejected
	CreatedBy  uint
	ApprovedBy *uint
	ApprovedAt *time.Time
	Items      []MaintenanceRequestItem
	Vehicle    Vehicle
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
