package main

import (
	"myjwtauthentication/database"
	"myjwtauthentication/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Database connection
	database.ConnectDB()
	app := fiber.New()

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	// routes
	routes.RouteSetup(app)

	app.Listen(":8000")
}
