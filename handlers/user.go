package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/models"
	service "github.com/gotired/golang_backend/services"
)

func ListUsers(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/ListUsers"))
	}
	if users == nil {
		users = make([]models.User, 0)
	}
	return c.JSON(Success{}.Data(users))
}
