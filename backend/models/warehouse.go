package models

import "time"

// --- Warehouse ---
type Warehouse struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Location  string
	Stocks    []Stock
	CreatedAt time.Time
	UpdatedAt time.Time
}
