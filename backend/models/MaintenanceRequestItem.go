package models

type MaintenanceRequestItem struct {
	ID        uint `gorm:"primaryKey"`
	RequestID uint
	ItemID    uint
	Quantity  float64
}
