package model

import "time"

type Product struct {
	Id          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title" form:"title"`
	Description string     `json:"description"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP;type:timestamp;not null" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"type:timestamp" json:"updated_at"`
}
