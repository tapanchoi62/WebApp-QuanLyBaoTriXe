package models

import (
	"time"

	"gorm.io/gorm"
)

// --- Role ---
type Role struct {
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"unique;not null"`            // ví dụ: Admin, Technician, Warehouse
	Permissions []Permission `gorm:"many2many:role_permissions"` // quan hệ many-to-many
	Users       []User
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
