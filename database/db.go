package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gotired/golang_backend/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	dbURL := "postgres://" + cfg.DBUsername + ":" + cfg.DBPassword +
		"@" + cfg.DBHost + ":" + strconv.Itoa(cfg.DBPort) +
		"/" + cfg.DBName + "?sslmode=disable"

	fmt.Println(dbURL)
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
