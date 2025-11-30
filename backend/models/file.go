package models

import "time"

type File struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	EntityID   string    `json:"entityId"`
	EntityType string    `json:"entityType"`
	FileURL    string    `json:"fileUrl"`
	FileName   string    `json:"fileName"`
	MimeType   string    `json:"mimeType"`
	Size       int64     `json:"size"`
	CreatedAt  time.Time `json:"createdAt"`
}
