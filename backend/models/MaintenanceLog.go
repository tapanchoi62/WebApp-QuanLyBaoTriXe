package models

type MaintenanceLog struct {
	ID        uint `gorm:"primaryKey"`
	VehicleID uint
	Km        int
	Desc      string
}
