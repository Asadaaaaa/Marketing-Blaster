package main

import (
	"Marketing-Blaster/models"
	"Marketing-Blaster/models/migrations"
	"Marketing-Blaster/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	models.InitDatabase()
	migrations.RunMigration()

	routes.RouteHandler(app)

	app.Listen(":3000")
}
