package models

import (
	"time"

	"gorm.io/gorm"
)

// --- Permission ---
type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"` // ví dụ: CREATE_REQUEST, APPROVE_REQUEST...
	Roles     []Role `gorm:"many2many:role_permissions"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
