package models

import "time"

// --- StockLog (Nhật ký kho) ---
type StockLog struct {
	ID         uint    `gorm:"primaryKey"`
	StockID    uint    `gorm:"not null"`
	Type       string  `gorm:"not null"` // IN, OUT, ADJUST
	Quantity   float64 `gorm:"not null"`
	Note       string
	SupplierID *uint // nếu có
	Stock      Stock
	Supplier   *Supplier
	CreatedAt  time.Time
}
