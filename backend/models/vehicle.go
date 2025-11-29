package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model    `swaggerignore:"true"`
	Plate_Number  string `gorm:"unique" json:"plate_number"`
	Model_Vehicle string `json:"model_vehicle"`
	Year          int    `json:"year"`
	Note          string `json:"note"`
}
