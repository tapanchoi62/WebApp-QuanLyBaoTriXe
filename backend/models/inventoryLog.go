package models

type InventoryLog struct {
	ID        uint `gorm:"primaryKey"`
	ItemID    uint
	Type      string // IN / OUT / ADJUST
	Quantity  float64
	CreatedBy uint
}
