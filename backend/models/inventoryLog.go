package models

import "gorm.io/gorm"

type InventoryLog struct {
	gorm.Model  `swaggerignore:"true"`
	InventoryID uint   `json:"inventory_id"`
	Action      string `json:"action"` // IN | OUT | ADJUST
	Quantity    int    `json:"quantity"`
	Note        string `json:"note"`
}
