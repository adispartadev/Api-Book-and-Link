package services

import (
	"api.go/database"
	"api.go/entity"
	"api.go/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

func GetUserLogin(c *fiber.Ctx, tokenType string) (model.User, bool) {
	var UserLogin model.User

	tokenString := c.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	if tokenString == "" {
		return UserLogin, false
	}
	var db = database.GetDbInstance()

	//check token was in blacklist token
	var blackListToken model.BlackListToken
	db.Where("token = ?", tokenString).First(&blackListToken)

	if blackListToken.Id != 0 {
		return UserLogin, false
	}

	// check token
	parseStatus, claimResult, _ := ParseJwtToken(tokenString)

	if parseStatus == false {
		return UserLogin, false
	}

	// token type
	if claimResult.TokenType != tokenType {
		return UserLogin, false
	}

	IdUser := claimResult.Id
	db.Where("id = ?", IdUser).First(&UserLogin)

	return UserLogin, true
}

func CreateJwtToken(user model.User, expirationTime time.Time, tokenType string) (tokenString string, status bool, error any) {
	// created JWT Key
	secret := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(secret)

	claims := &entity.JwtClaims{
		Id:          int(user.Id),
		User:        user,
		GeneratedAt: time.Now(),
		TokenType:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", false, err.Error()
	}

	return tokenString, true, nil

}

func ParseJwtToken(tokenString string) (status bool, claim *entity.JwtClaims, data any) {
	// check token
	secret := os.Getenv("JWT_SECRET")
	var jwtKey = []byte(secret)

	claims := &entity.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false, claims, err
	}

	claimResult, _ := token.Claims.(*entity.JwtClaims)
	return true, claimResult, nil
}
