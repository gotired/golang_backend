package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models"
	service "github.com/gotired/golang_backend/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type User interface {
// 	List(c *fiber.Ctx) error
// 	Register(c *fiber.Ctx) error
// 	Login(c *fiber.Ctx) error
// }

type UserStruct struct {
}

func (u *UserStruct) List(c *fiber.Ctx) error {
	users, err := service.GetAllUsers()
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

	db := database.GetDB()

	var existingUser models.User
	if err := db.Where("email = ?", requestBody.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			Failure{}.Detail("Email is already registered", "handlers/user/Register"))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// If an unexpected error occurred during the query, return it
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userID := uuid.New()

	user := models.User{ID: userID.String(), Email: requestBody.Email, Password: string(hashedPassword), FirstName: requestBody.FirstName, LastName: requestBody.LastName, Phone: requestBody.Phone, Username: requestBody.Username}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(Success{}.Detail("User registered successfully!"))
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

	// Passwords match, login successful
	return c.SendString("Login successful!")
}
