package model

import "time"

type User struct {
	Id        uint       `gorm:"primaryKey" json:"id"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:current_timestamp" json:"updated_at"`
}
