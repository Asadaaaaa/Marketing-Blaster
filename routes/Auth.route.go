package routes

import (
	"Marketing-Blaster/controllers"

	"github.com/gofiber/fiber/v2"
)

func InitAuth(app *fiber.App) {
	routePrefix := "/auth"

	RouteAuth(app, routePrefix)
}

func RouteAuth(app *fiber.App, routePrefix string) {
	// Register Route
	app.Post(routePrefix+"/register", controllers.RegisterAuthController)
	app.Post(routePrefix+"/login", controllers.LoginAuthController)
}
