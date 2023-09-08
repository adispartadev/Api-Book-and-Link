package services

import (
	"api.go/database"
	"api.go/entity"
	"api.go/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
)

func GetUserLogin(c *fiber.Ctx) (model.User, bool) {
	var UserLogin model.User

	tokenString := c.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	if tokenString == "" {
		return UserLogin, false
	}

	// check token
	secret := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(secret)

	claims := &entity.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return UserLogin, false
	}

	claimResult, _ := token.Claims.(*entity.JwtClaims)
	IdUser := claimResult.Id

	var db = database.GetDbInstance()
	db.Where("id = ?", IdUser).First(&UserLogin)

	return UserLogin, true
}
