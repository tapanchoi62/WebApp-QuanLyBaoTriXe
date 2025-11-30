package models

type MaintenanceRequestItem struct {
	ID                   uint    `gorm:"primaryKey"`
	MaintenanceRequestID uint    `gorm:"not null"`
	ItemID               uint    `gorm:"not null"`
	Quantity             float64 `gorm:"not null"`
	Item                 Item
}
