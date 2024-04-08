package database

import (
	"log"
	"strconv"

	"gorm.io/driver/postgres"

	"github.com/gotired/golang_backend/config"
	Table "github.com/gotired/golang_backend/models/table"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var db *gorm.DB

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
	db.AutoMigrate(&Table.Role{}, &Table.UserInfo{}, &Table.UserCredential{})

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
