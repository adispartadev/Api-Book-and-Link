package routes

import (
	apiController "api.go/controllers/api"
	"api.go/middleware"
	"github.com/gofiber/fiber/v2"
)

func AppApiRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/", apiController.HelloWorld)
	api.Get("/profile", middleware.AuthApi(), apiController.UserLogin)
	api.Post("/register", apiController.RegisterUser)
	api.Post("/login", apiController.LoginUser)
	api.Post("/forgot-password", apiController.ForgotPassword)
	api.Post("/reset-password", apiController.ResetPassword)
}
