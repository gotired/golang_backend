package models

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Email     string `gorm:"uniqueIndex"`
	Username  string `gorm:"uniqueIndex"`
	FirstName string
	LastName  string
	Phone     string
	Password  string
}
