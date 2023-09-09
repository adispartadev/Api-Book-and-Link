package api

import (
	"api.go/database"
	"api.go/entity"
	"api.go/model"
	"api.go/services"
	"github.com/gofiber/fiber/v2"
)

func AllProduct(c *fiber.Ctx) error {
	var db = database.GetDbInstance()
	var products []model.Product
	db.Find(&products)

	return services.ApiJsonResponse(c, entity.Success, "Product List", products)
}

func DetailProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var db = database.GetDbInstance()
	var product model.Product
	db.First(&product, id)

	return services.ApiJsonResponse(c, entity.Success, "Product Detail", product)
}
