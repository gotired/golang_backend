package service

import (
	"github.com/gotired/golang_backend/database"
	"github.com/gotired/golang_backend/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	db := database.GetDB()

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
