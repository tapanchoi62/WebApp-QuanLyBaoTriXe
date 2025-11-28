package models

type MaintenanceRequest struct {
	ID        uint `gorm:"primaryKey"`
	VehicleID uint
	Status    string // Pending / Approved / Rejected
}
