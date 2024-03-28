package database

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gotired/golang_backend/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	dbURL := "postgres://" + cfg.DBUsername + ":" + cfg.DBPassword +
		"@" + cfg.DBHost + ":" + strconv.Itoa(cfg.DBPort) +
		"/" + cfg.DBName + "?sslmode=disable"

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
