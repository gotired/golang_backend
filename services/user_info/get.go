package services

import (
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models/table"
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
