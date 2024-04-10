package utils

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gotired/golang_backend/config"
)

func GenerateJWT(userID string, expiresIn time.Duration) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expiresIn).Unix(),
	})

	signedToken, err := token.SignedString([]byte(cfg.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
