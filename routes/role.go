package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/handlers"
)

func SetupRoleRoutes(app *fiber.App) {
	roleRoutes := app.Group("/roles")
	roleHandler := &handlers.RoleStruct{}

	roleRoutes.Get("/", roleHandler.List)
	roleRoutes.Post("/register", roleHandler.Register)
}
