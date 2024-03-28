package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	APPPort    int
}

func Load() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, errors.New("error loading .env: " + err.Error())
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return Config{}, errors.New("DB_PORT value error: " + err.Error())
	}

	appPortStr := os.Getenv("APP_PORT")
	appPort, err := strconv.Atoi(appPortStr)
	if err != nil {
		return Config{}, errors.New("APP_PORT value error: " + err.Error())
	}

	return Config{
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
