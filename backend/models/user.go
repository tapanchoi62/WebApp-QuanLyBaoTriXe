package models

import (
	"time"

	"gorm.io/gorm"
)

// --- User ---
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	RoleID       uint   `gorm:"not null"`
	Role         Role
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	TokenVersion uint           // tăng khi đổi role/permission

}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}
