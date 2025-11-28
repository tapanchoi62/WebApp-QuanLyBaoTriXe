package models

import (
	"time"
)

type Vehicle struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Plate_Number string    `gorm:"unique" json:"plate_number"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Note         string    `json:"note"`
	CreatedAt    time.Time `json:"created_at"`
}
