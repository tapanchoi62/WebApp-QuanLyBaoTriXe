package models

import "time"

type Warehouse struct {
	ID        uint      `gorm:"primaryKey" json:"ID"`
	Name      string    `gorm:"not null" json:"name"`
	Location  string    `json:"location"`
	Stocks    []Stock   `json:"-"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
