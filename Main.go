package main

import (
	"Marketing-Blaster/models"
	"Marketing-Blaster/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	app := fiber.New()

	models.InitDatabase()
	models.RunMigration()

	routes.RouteHandler(app)

	app.Listen(":3000")
}
