package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/users")
	userHandler := &handlers.User{}

	userRoutes.Get("/", userHandler.List)
	userRoutes.Post("/", userHandler.Register)
}
