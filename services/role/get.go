package services

import (
	"errors"

	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models/table"
	"gorm.io/gorm"
)

func Validate(role string) error {
	db := database.GetDB()

	var existingRole table.Role
	if err := db.Where("id = ?", role).First(&existingRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("role does not exist")
		}
		return err
	}

	return nil
}
