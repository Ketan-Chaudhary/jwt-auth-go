package routes

import (
	"myjwtauthentication/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteSetup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Post("/api/register", controllers.RegisterUser)
	app.Post("/api/login", controllers.LoginUser)
	app.Get("/api/user", controllers.User)
	app.Post("/api/user", controllers.Logout)
}
