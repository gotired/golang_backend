package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/models/table"
	roleService "github.com/gotired/golang_backend/services/role"
)

type RoleStruct struct {
}

func (u *RoleStruct) List(c *fiber.Ctx) error {
	roles, err := roleService.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/role/List"))
	}
	if roles == nil {
		roles = make([]table.Role, 0)
	}

	return c.JSON(Success{}.Data(roles))
}

func (u *RoleStruct) Register(c *fiber.Ctx) error {
	var requestBody table.Role
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			Failure{}.Detail(err.Error(), "handlers/role/Register"))
	}

	roleID, err := roleService.Insert(requestBody.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/role/Register"))
	}
	return c.JSON(Success{}.Detail("Role registered successfully! " + (*roleID).String()))
}
