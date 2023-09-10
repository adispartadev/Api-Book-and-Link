package entity

import (
	"api.go/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtClaims struct {
	Id          int `json:"id"`
	User        model.User
	GeneratedAt time.Time
	TokenType   string
	jwt.RegisteredClaims
}
