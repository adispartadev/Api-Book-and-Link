package middleware

import (
	"api.go/entity"
	"api.go/services"
	"github.com/gofiber/fiber/v2"
)

func AuthApi() fiber.Handler {

	return func(c *fiber.Ctx) error {
		UserLogin, isTrue := services.GetUserLogin(c, entity.JwtToken)
		if !isTrue {
			return services.ApiJsonResponseWithCode(c, entity.Error, "Please login to continue", nil, 401)
		}

		if UserLogin.Id == 0 {
			return services.ApiJsonResponseWithCode(c, entity.Error, "User not found, please sign in", nil, 401)
		}

		return c.Next()
	}
}

func AuthApiRefresh() fiber.Handler {
	return func(c *fiber.Ctx) error {
		UserLogin, isTrue := services.GetUserLogin(c, entity.JwtRefresh)
		if !isTrue {
			return services.ApiJsonResponseWithCode(c, entity.Error, "Please login to continue", nil, 401)
		}

		if UserLogin.Id == 0 {
			return services.ApiJsonResponseWithCode(c, entity.Error, "User not found, please sign in", nil, 401)
		}

		return c.Next()
	}
}
