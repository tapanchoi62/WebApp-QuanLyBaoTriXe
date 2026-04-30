package models

import "time"

type Stock struct {
	ID          uint       `gorm:"primaryKey" json:"ID"`
	ItemID      uint       `gorm:"not null" json:"item_id"`
	WarehouseID uint       `gorm:"not null" json:"warehouse_id"`
	Quantity    float64    `gorm:"default:0" json:"quantity"`
	Item        Item       `json:"Item"`
	Warehouse   Warehouse  `json:"Warehouse"`
	StockLogs   []StockLog `json:"-"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	UpdatedAt   time.Time  `json:"UpdatedAt"`
}
