package bootstrap

import (
	"api.go/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
)

func BootApplication() {

	err := godotenv.Load()
	if err != nil {
		panic("Unable to load .env file")
	}

	// new fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Accept-Language, Content-Length",
	}))

	// registering routes
	routes.AppApiRoutes(app)

	// listenig port
	fmt.Println("Listening to :" + os.Getenv("PORT"))
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
