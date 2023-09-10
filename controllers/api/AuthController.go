package api

import (
	"api.go/database"
	"api.go/entity"
	"api.go/model"
	"api.go/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

func RegisterUser(c *fiber.Ctx) error {

	// validation rule
	type FormDataStrct struct {
		FullName        string `json:"full_name" form:"full_name" validate:"required"`
		Email           string `json:"email" form:"email" validate:"required"`
		Password        string `json:"password" form:"password" validate:"required"`
		PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required,eqfield=Password"`
	}

	// validating requet
	var formData = FormDataStrct{}
	validationStatus, errorField, message := services.ValidatingRequest(c, &formData)
	if validationStatus == false {
		return services.ApiJsonResponse(c, entity.Error, message, errorField)
	}

	var db = database.GetDbInstance()

	// validating email
	var userExists model.User
	db.Where("email = ?", formData.Email).First(&userExists)
	if userExists.Id != 0 {
		return services.ApiJsonResponse(c, entity.Error, "Email is already used", nil)
	}

	// saving user
	user := model.User{}
	user.FullName = formData.FullName
	user.Email = formData.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}
	user.Password = string(hashedPassword)
	db.Save(&user)

	return services.ApiJsonResponse(c, entity.Success, "Registered successfully", nil)
}

// Login A User
func LoginUser(c *fiber.Ctx) error {
	// validation rule
	type FormDataStrct struct {
		Email    string `json:"email" form:"email" validate:"required"`
		Password string `json:"password" form:"password" validate:"required"`
	}

	// validating requet
	var formData = FormDataStrct{}
	validationStatus, errorField, message := services.ValidatingRequest(c, &formData)
	if validationStatus == false {
		return services.ApiJsonResponse(c, entity.Error, message, errorField)
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
	expirationTime := time.Now().Add(time.Hour * 24)
	tokenString, status, errr := services.CreateJwtToken(user, expirationTime, entity.JwtToken)
	if status == false {
		return services.ApiJsonResponse(c, entity.Error, "Unable to generate access token", errr)
	}

	// create refresh token
	expirationTime = time.Now().Add(time.Hour * 24 * 30) // 1 month
	refreshTokenString, status, errr := services.CreateJwtToken(user, expirationTime, entity.JwtRefresh)
	if status == false {
		return services.ApiJsonResponse(c, entity.Error, "Unable to generate refresh token", errr)
	}

	return services.ApiJsonResponse(c, entity.Success, "Login successfully", map[string]any{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
		"user":          user,
	})
}

func ForgotPassword(c *fiber.Ctx) error {

	// validation rule
	type FormDataStrct struct {
		Email string `json:"email" form:"email" validate:"required"`
	}

	// validating requet
	var formData = FormDataStrct{}
	validationStatus, errorField, message := services.ValidatingRequest(c, &formData)
	if validationStatus == false {
		return services.ApiJsonResponse(c, entity.Error, message, errorField)
	}

	// get user by email
	var user model.User
	var db = database.GetDbInstance()
	db.Where("email = ?", formData.Email).First(&user)

	if user.Id == 0 {
		return services.ApiJsonResponse(c, entity.Error, "User not found", nil)
	}

	var randString = services.RandSeq(10)
	user.PasswordToken = randString
	db.Save(&user)

	var link = os.Getenv("FRONT_PAGE_BASE_URL") + "/auth/reset-password/" + randString
	err := services.SendEmail(user.Email, "Forgot Password", "Please open "+link+" to reset your password.")
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}
	return services.ApiJsonResponse(c, entity.Success, "Please open your email for password recovery", nil)
}

func ResetPassword(c *fiber.Ctx) error {
	// validation rule
	type FormDataStrct struct {
		Password        string `json:"password" form:"password" validate:"required"`
		PasswordConfirm string `json:"password_confirm" form:"password_confirm" validate:"required,eqfield=Password"`
		Token           string `json:"token" form:"token" validate:"required"`
	}

	// validating requet
	var formData = FormDataStrct{}
	validationStatus, errorField, message := services.ValidatingRequest(c, &formData)
	if validationStatus == false {
		return services.ApiJsonResponse(c, entity.Error, message, errorField)
	}

	// get user by email
	var user model.User
	var db = database.GetDbInstance()
	db.Where("password_token = ?", formData.Token).First(&user)

	if user.Id == 0 {
		return services.ApiJsonResponse(c, entity.Error, "User not found or Url expired", nil)
	}

	// saving user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), 14)
	if err != nil {
		return services.ApiJsonResponse(c, entity.Error, "Error occurred", err.Error())
	}
	user.Password = string(hashedPassword)
	db.Save(&user)
	return services.ApiJsonResponse(c, entity.Success, "Your password successfully change. Please login to continue", nil)
}

func UserLogin(c *fiber.Ctx) error {
	UserLogin, _ := services.GetUserLogin(c, entity.JwtToken)
	return services.ApiJsonResponse(c, entity.Success, "Logged in user", UserLogin)
}

func LogoutUser(c *fiber.Ctx) error {

	// validation rule
	type FormDataStrct struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
	}

	// validating requet
	var formData = FormDataStrct{}
	validationStatus, errorField, message := services.ValidatingRequest(c, &formData)
	if validationStatus == false {
		return services.ApiJsonResponse(c, entity.Error, message, errorField)
	}

	var db = database.GetDbInstance()

	// add token to blacklist data
	tokenString := c.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	var blackListToken model.BlackListToken
	blackListToken.Token = tokenString
	db.Save(&blackListToken)

	// add refresh token to blacklist data
	var blackListRefreshToken model.BlackListToken
	tokenRString := formData.RefreshToken
	tokenRString = strings.Replace(tokenRString, "Bearer ", "", 1)
	blackListRefreshToken.Token = tokenRString
	db.Save(&blackListRefreshToken)

	return services.ApiJsonResponse(c, entity.Success, "Logout success", nil)
}

func RefreshUserToken(c *fiber.Ctx) error {
	user, _ := services.GetUserLogin(c, entity.JwtRefresh)

	// created JWT Key
	expirationTime := time.Now().Add(time.Hour * 24)
	tokenString, status, errr := services.CreateJwtToken(user, expirationTime, entity.JwtToken)
	if status == false {
		return services.ApiJsonResponse(c, entity.Error, "Unable to generate access token", errr)
	}

	// create refresh token
	expirationTime = time.Now().Add(time.Hour * 24 * 30) // 1 month
	refreshTokenString, status, errr := services.CreateJwtToken(user, expirationTime, entity.JwtRefresh)
	if status == false {
		return services.ApiJsonResponse(c, entity.Error, "Unable to generate refresh token", errr)
	}

	return services.ApiJsonResponse(c, entity.Success, "Login successfully", map[string]any{
		"token":         tokenString,
		"refresh_token": refreshTokenString,
		"user":          user,
	})

}
