package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/gotired/golang_backend/models"
	"github.com/joho/godotenv"
)

func Load() (models.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return models.Config{}, errors.New("error loading .env: " + err.Error())
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return models.Config{}, errors.New("DB_PORT value error: " + err.Error())
	}

	appPortStr := os.Getenv("APP_PORT")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return models.Config{}, errors.New("APP_PORT value error: " + err.Error())
	}

	return models.Config{
		DBUsername: getEnv("DB_USERNAME", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     dbPort,
		DBName:     getEnv("DB_NAME", ""),
		APPPort:    appPort,
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
