package api

import (
	"api.go/database"
	"api.go/entity"
	"api.go/model"
	"api.go/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// Register User
func RegisterUser(c *fiber.Ctx) error {

	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}

	var db = database.GetDbInstance()
	user.Password = string(hashedPassword)
	db.Save(&user)
	return services.ApiJsonResponse(c, entity.Success, "Registered successfully", nil)
}

// Login A User
func LoginUser(c *fiber.Ctx) error {
	formData := model.User{}
	if err := c.BodyParser(&formData); err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}

	// get user by email
	var user model.User
	var db = database.GetDbInstance()
	db.Where("email = ?", formData.Email).First(&user)

	if user.Id == 0 {
		return services.ApiJsonResponse(c, entity.Error, "User not found", nil)
	}

	// check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Wrong password and email combination", nil)
	}

	// created JWT Key
	secret := os.Getenv("JWT_SECRET")
	expirationTime := time.Now().Add(60 * 24 * 30 * 6 * time.Minute)
	var jwtKey = []byte(secret)

	claims := &entity.JwtClaims{
		IdUser: int(user.Id),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}

	return services.ApiJsonResponse(c, entity.Success, "Login successfully", map[string]any{
		"token": tokenString,
		"user":  user,
	})
}
