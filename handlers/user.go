package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/models"
	"github.com/gotired/golang_backend/models/table"
	user_credential_service "github.com/gotired/golang_backend/services/user_credential"
	user_info_service "github.com/gotired/golang_backend/services/user_info"
	"github.com/gotired/golang_backend/utils"
)

type UserStruct struct {
}

func (u *UserStruct) List(c *fiber.Ctx) error {
	users, err := user_info_service.Get()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/List"))
	}
	if users == nil {
		users = make([]table.UserInfo, 0)
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

	user_id, err := user_credential_service.Validate(loginData.Identifier, loginData.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Login"))
	}
	accessToken, err := utils.GenerateJWT((*user_id).String(), time.Minute*15)
	if err != nil {
		return err
	}

	refreshToken, err := utils.GenerateJWT((*user_id).String(), time.Hour)
	if err != nil {
		return err
	}
	response := make(map[string]string)
	response["access_token"] = accessToken
	response["refresh_token"] = refreshToken

	return c.JSON(Success{}.Data(response))
}
