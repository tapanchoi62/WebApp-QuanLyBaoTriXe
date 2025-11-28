package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `json:"password"`
	Role      string    `gorm:"type:enum('admin','technician','warehouse');default:'technician'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
