package models

type Inventory struct {
	ID        uint `gorm:"primaryKey"`
	Warehouse string
	ItemID    uint
	Quantity  float64
}
