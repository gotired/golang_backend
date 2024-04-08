package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models/table"
	role_service "github.com/gotired/golang_backend/services/role"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Validate(Identifier string, password string) error {
	db := database.GetDB()

	var user_credential table.UserCredential
	if err := db.Where("user_name = ? OR email = ?", Identifier, Identifier).First(&user_credential).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user_credential.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}

func Register(email string, username string, role string, password string) (*uuid.UUID, error) {
	db := database.GetDB()

	var existingUser table.UserCredential
	if err := db.Where("email = ? OR user_name = ?", email, username).First(&existingUser).Error; err == nil {
		return nil, err
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	err := role_service.Validate(role)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := table.UserCredential{Email: email, Password: string(hashedPassword), UserName: username, Role: role}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, err
}
