package routes

import (
	"Marketing-Blaster/controllers"
	"Marketing-Blaster/middlewares"

	"github.com/gofiber/fiber/v2"
)

func InitAdmin(app *fiber.App) {
	routePrefix := "/admin"

	RouteAdmin(app, routePrefix)
}

func RouteAdmin(app *fiber.App, routePrefix string) {
	app.Post(routePrefix+"/auth/login", controllers.AdminLoginAuthController)
	// get all user data
	app.Get(routePrefix+"/user/get", middlewares.AuthorizationMiddlewareAdmin, controllers.AdminGetAllUserController)
	// get all mail logs
	app.Get(routePrefix+"/mail/get", middlewares.AuthorizationMiddlewareAdmin, controllers.AdminGetAllMailLogsController)

}
