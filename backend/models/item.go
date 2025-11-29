package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model `swaggerignore:"true"`
	Name       string
	Unit       string
	Category   string
}
