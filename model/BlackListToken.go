package model

import "time"

type BlackListToken struct {
	Id        uint       `gorm:"primaryKey" json:"id"`
	Token     string     `json:"token" form:"token"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP;type:timestamp;not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updated_at"`
}
