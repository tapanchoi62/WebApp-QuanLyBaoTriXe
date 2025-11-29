package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"unique;not null" json:"username"`
	Password   string `json:"password"`
	Role       string `gorm:"type:enum('admin','technician','warehouse');default:'technician'" json:"role"`
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
