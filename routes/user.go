package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")
	userRoutes.Get("/", handlers.ListUsers)
}
