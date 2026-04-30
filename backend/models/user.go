package models

import (
	"time"

	"gorm.io/gorm"
)

// --- User ---
type User struct {
	ID           uint           `gorm:"primaryKey" json:"ID"`
	Username     string         `gorm:"unique;not null" json:"Username"`
	Password     string         `gorm:"not null" json:"-"`
	RoleID       uint           `gorm:"not null" json:"RoleID"`
	Role         Role           `json:"Role"`
	CreatedAt    time.Time      `json:"CreatedAt"`
	UpdatedAt    time.Time      `json:"UpdatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	TokenVersion uint           `json:"-"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}
