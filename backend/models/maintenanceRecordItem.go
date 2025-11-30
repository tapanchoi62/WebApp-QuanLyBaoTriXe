package models

// --- MaintenanceRecordItem (Vật tư sử dụng) ---
type MaintenanceRecordItem struct {
	ID                  uint    `gorm:"primaryKey"`
	MaintenanceRecordID uint    `gorm:"not null"`
	ItemID              uint    `gorm:"not null"`
	QuantityUsed        float64 `gorm:"not null"`
	Item                Item
}
