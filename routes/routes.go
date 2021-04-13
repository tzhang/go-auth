package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tzhang/go-auth/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/api/register", controllers.Register)

}
