package models

import "time"

// --- MaintenanceRecord (Nhật ký bảo trì) ---
type MaintenanceRecord struct {
	ID                   uint  `gorm:"primaryKey"`
	VehicleID            uint  `gorm:"not null"`
	MaintenanceRequestID *uint // nếu liên kết phiếu
	Description          string
	KM                   float64
	StartTime            time.Time
	EndTime              time.Time
	CreatedBy            uint
	Items                []MaintenanceRecordItem
	Vehicle              Vehicle
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
