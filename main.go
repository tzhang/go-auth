package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tzhang/go_auth/database"
	"github.com/tzhang/go_auth/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
