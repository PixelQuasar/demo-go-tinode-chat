package utils

import (
	"demo-go-tinode-chat/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
