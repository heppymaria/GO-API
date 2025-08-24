package routes

import (
	"go-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Get("/users/:id", controllers.GetUserByID)
}

