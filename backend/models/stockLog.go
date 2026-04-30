package models

import "time"

type StockLog struct {
	ID         uint      `gorm:"primaryKey" json:"ID"`
	StockID    uint      `gorm:"not null" json:"stock_id"`
	Type       string    `gorm:"not null" json:"type"`
	Quantity   float64   `gorm:"not null" json:"quantity"`
	Note       string    `json:"note"`
	SupplierID *uint     `json:"supplier_id"`
	Stock      Stock     `json:"-"`
	Supplier   *Supplier `json:"Supplier,omitempty"`
	CreatedAt  time.Time `json:"CreatedAt"`
}
