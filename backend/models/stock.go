package models

import "time"

// --- Stock (Tồn kho) ---
type Stock struct {
	ID          uint    `gorm:"primaryKey"`
	ItemID      uint    `gorm:"not null"`
	WarehouseID uint    `gorm:"not null"`
	Quantity    float64 `gorm:"default:0"`
	Item        Item
	Warehouse   Warehouse
	StockLogs   []StockLog
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
