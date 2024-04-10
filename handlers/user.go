package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/golang_backend/models"
	"github.com/gotired/golang_backend/models/table"
	userCredentialService "github.com/gotired/golang_backend/services/user_credential"
	userInfoService "github.com/gotired/golang_backend/services/user_info"
	"github.com/gotired/golang_backend/utils"
)

type UserStruct struct {
}

func (u *UserStruct) List(c *fiber.Ctx) error {
	users, err := userInfoService.Get()
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

	userID, err := userCredentialService.Insert(requestBody.Email, requestBody.Username, requestBody.Role, requestBody.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Register"))
	}

	err = userInfoService.Insert(*userID, requestBody.FirstName, requestBody.LastName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Register"))
	}

	return c.JSON(Success{}.Detail("User registered successfully! "))
}

func (u *UserStruct) Login(c *fiber.Ctx) error {
	var loginData models.UserLogin
	if err := c.BodyParser(&loginData); err != nil {
		return err
	}

	userID, err := userCredentialService.Validate(loginData.Identifier, loginData.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			Failure{}.Detail(err.Error(), "handlers/user/Login"))
	}
	accessToken, err := utils.GenerateJWT((*userID).String(), time.Minute*15)
	if err != nil {
		return err
	}

	refreshToken, err := utils.GenerateJWT((*userID).String(), time.Hour)
	if err != nil {
		return err
	}
	response := make(map[string]string)
	response["access_token"] = accessToken
	response["refresh_token"] = refreshToken

	return c.JSON(Success{}.Data(response))
}
