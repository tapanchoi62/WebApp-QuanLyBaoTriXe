package models

import "time"

// --- Supplier (Nhà cung cấp) ---
type Supplier struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Contact   string
	StockLogs []StockLog
	CreatedAt time.Time
	UpdatedAt time.Time
}
