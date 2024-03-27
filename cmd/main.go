package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/routes"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	routes.SetupUserRoutes(app)

	// Start the server on port 3000
	app.Listen(":3000")
}
