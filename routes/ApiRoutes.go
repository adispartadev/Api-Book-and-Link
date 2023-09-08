package routes

import (
	apiController "api.go/controllers/api"
	"github.com/gofiber/fiber/v2"
)

func AppApiRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/", apiController.HelloWorld)

}
