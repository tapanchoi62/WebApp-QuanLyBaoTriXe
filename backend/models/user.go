package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Role      string    `gorm:"type:enum('admin','technician','warehouse');default:'technician'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
