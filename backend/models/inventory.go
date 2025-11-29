package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model  `swaggerignore:"true"`
	Warehouse   string `json:"warehouse"`             // Tên kho
	ItemCode    string `json:"itemCode"`              // Mã vật tư
	ItemName    string `json:"itemName"`              // Tên vật tư
	Quantity    int    `json:"quantity"`              // Tồn kho
	Unit        string `json:"unit"`                  // Đơn vị tính
	MinStock    int    `json:"minStock"`              // Mức tồn kho cảnh báo
	Description string `json:"description,omitempty"` // Ghi chú
}
