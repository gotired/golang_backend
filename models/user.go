package models

type UserDetail struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"num"`
	Email     string `json:"email"`
}

type UserRegister struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"user_name"`
	Password  string `json:"password"`
	Confirm   string `json:"confirm"`
	Phone     string `json:"num"`
}

var users []UserDetail
