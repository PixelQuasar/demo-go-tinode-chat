package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type Claims struct {
	Username string `bson:"username"`
	jwt.RegisteredClaims
}
