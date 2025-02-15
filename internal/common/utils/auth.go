package utils

import (
	"demo-go-tinode-chat/config"
	"demo-go-tinode-chat/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(config.AppConfig.JwtExpirationMinutes))
	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.AppConfig.JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUsernameByToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return config.AppConfig.JwtKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	username := claims["Username"].(string)
	return username, nil
}
