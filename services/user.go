package service

import (
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models"
)

func GetAllUsers() ([]models.UserDetail, error) {
	var users []models.UserDetail
	db := database.GetDB()

	if err := db.Table("users").Select("users.*").
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
