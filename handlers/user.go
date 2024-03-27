package handlers

import (
	"github.com/gofiber/fiber/v2"
	service "github.com/gotired/golang_backend/services"
)

func ListUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(users)
}
