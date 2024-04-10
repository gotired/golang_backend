package models

type UserBase struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"num"`
}

type UserDetail struct {
	UserBase
	ID string `json:"id" gorm:"primaryKey"`
}

type UserRegister struct {
	UserBase
	Username string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Confirm  string `json:"confirm" validate:"required,eqfield=Password"`
	Role     string `json:"role" validate:"required,uuid"`
}

type UserLogin struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

var users []UserDetail
