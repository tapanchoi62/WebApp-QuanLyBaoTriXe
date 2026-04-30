package models

import "time"

type Item struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	Name      string    `gorm:"not null" json:"name"`
	Category  string    `json:"category"`
	Unit      string    `json:"unit"`
	Stocks    []Stock   `json:"-"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
