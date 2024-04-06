package database

import (
	"log"
	"strconv"

	"gorm.io/driver/postgres"

	"github.com/gotired/golang_backend/config"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var db *gorm.DB

type Users struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"num"`
	Username  string `json:"user_name" validate:"required"`
	Password  string `json:"password" validate:"required,min=6"`
}

func init() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUsername + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + strconv.Itoa(cfg.DBPort) + " sslmode=disable TimeZone=Asia/Bangkok"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbConnection
	db.AutoMigrate(&Users{})
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
