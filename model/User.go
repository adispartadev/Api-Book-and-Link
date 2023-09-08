package model

import "time"

type User struct {
	Id            uint       `gorm:"primaryKey" json:"id"`
	FullName      string     `json:"full_name" form:"full_name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	PasswordToken string     `json:"password_token"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP;type:timestamp;not null" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"type:timestamp" json:"updated_at"`
}
