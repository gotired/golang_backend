package services

import (
	"errors"

	"github.com/google/uuid"
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

func Insert(role string) (*uuid.UUID, error) {
	db := database.GetDB()

	var existingRole table.Role
	if err := db.Where("role = ?", role).First(&existingRole).Error; err == nil {
		return nil, err
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	roleDetail := table.Role{Role: role}
	if err := db.Create(&roleDetail).Error; err != nil {
		return nil, err
	}
	return &roleDetail.ID, nil
}

func Get() ([]table.Role, error) {
	var roles []table.Role
	db := database.GetDB()

	if err := db.Table("roles").Select("*").
		Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}
