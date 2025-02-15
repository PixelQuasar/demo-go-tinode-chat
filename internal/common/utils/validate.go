package utils

import (
	"context"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIfUserExists(username string) (bool, error) {
	var result models.User
	err := db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasLetter := false
	for _, char := range password {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return errors.New("password must contain at least one letter")
	}

	hasDigit := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasDigit = true
			break
		}
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}

	return nil
}
