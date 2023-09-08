package services

import (
	"api.go/entity"
	"github.com/gofiber/fiber/v2"
)

func ApiJsonResponse(c *fiber.Ctx, status string, message string, data any) error {
	var result = entity.ApiFormat{Status: status, Message: message, Data: data}
	return c.JSON(result)
}

func ApiJsonResponseWithCode(c *fiber.Ctx, status string, message string, data any, code int) error {
	var result = entity.ApiFormat{Status: status, Message: message, Data: data}
	return c.Status(code).JSON(result)
}
