package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models"
	"github.com/gotired/golang_backend/services"
	user_credential_service "github.com/gotired/golang_backend/services/user_credential"
	"golang.org/x/crypto/bcrypt"
)

type UserStruct struct {
}

func (u *UserStruct) List(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/List"))
	}
	if users == nil {
		users = make([]models.UserDetail, 0)
	}
	return c.JSON(Success{}.Data(users))
}

func (u *UserStruct) Register(c *fiber.Ctx) error {
	var requestBody models.UserRegister
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Register"))
	}
	if requestBody.Password != requestBody.Confirm {
		return c.Status(fiber.StatusBadRequest).JSON(
			Failure{}.Detail("Mismatch Password and Confirm password", "handlers/user/Register"))
	}

	user_id, err := user_credential_service.Register(requestBody.Email, requestBody.Username, requestBody.Username, requestBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Register"))
	}
	return c.JSON(Success{}.Detail("User registered successfully! " + (*user_id).String()))
}

func (u *UserStruct) Login(c *fiber.Ctx) error {
	var loginData models.UserLogin
	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	db := database.GetDB()

	var user models.UserRegister
	if err := db.Where("user_name = ? OR email = ?", loginData.Identifier, loginData.Identifier).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Login"))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Login"))
	}

	return c.JSON(Success{}.Detail("Login successful!"))
}
