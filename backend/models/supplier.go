package models

import "time"

type Supplier struct {
	ID        uint       `gorm:"primaryKey" json:"ID"`
	Name      string     `gorm:"not null" json:"name"`
	Contact   string     `json:"contact"`
	StockLogs []StockLog `json:"-"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
}
