package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tzhang/go-auth/database"
	"github.com/tzhang/go-auth/routes"
)

func main() {

	database.Connect()

	app := fiber.New()
	//加了cookie后，要增加以下这句
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	app.Listen(":8000")
}
