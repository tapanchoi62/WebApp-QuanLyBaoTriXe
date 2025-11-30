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

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Files    []File `json:"files"`
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
	FileID   string `json:"fileId"` // file đã upload trước
}

type UserInputCreate struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role"`
}
