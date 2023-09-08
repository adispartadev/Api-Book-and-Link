package api

import (
	"api.go/entity"
	"api.go/services"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return services.ApiJsonResponse(c, entity.Success, "Hello World", nil)
}
