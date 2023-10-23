package routes

import (
	"Marketing-Blaster/controllers"
	"Marketing-Blaster/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitMailer(app *fiber.App) {
	routePrefix := "/mailer"

	RouteMailer(app, routePrefix)
}

func RouteMailer(app *fiber.App, routePrefix string) {
	// Register Route
	app.Post(routePrefix+"/send", middlewares.AuthorizationMiddleware, controllers.SendMailerController)
	app.Post(routePrefix+"/ai/body", middlewares.AuthorizationMiddleware, controllers.AIGetBodyController)
	app.Post(routePrefix+"/ai/send", middlewares.AuthorizationMiddleware, controllers.AIGetBodyController)
}
