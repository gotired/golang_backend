package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/config"
	"github.com/gotired/golang_backend/routes"
	"github.com/gotired/golang_backend/utils"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	app := fiber.New()
	app.Use(utils.LogRequest)

	routes.SetupUserRoutes(app)

	app.Listen(":" + strconv.Itoa(cfg.APPPort))
}
