package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models/table"
	"gorm.io/gorm"
)

func Get() ([]table.UserInfo, error) {
	var users []table.UserInfo
	db := database.GetDB()

	if err := db.Table("user_infos").Select("*").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func Insert(userID uuid.UUID, first_name string, last_name string) error {
	db := database.GetDB()

	var existingUser table.UserInfo
	if err := db.Where("id = ?", userID).First(&existingUser).Error; err == nil {
		return err
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user := table.UserInfo{ID: userID, FirstName: first_name, LastName: last_name}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
