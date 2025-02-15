package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
