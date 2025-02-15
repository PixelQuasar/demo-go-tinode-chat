package services

import (
	"context"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	_, err = db.UserCollection.InsertOne(context.TODO(), user)
	return err
}

func AuthenticateUser(username, password string) (bool, error) {
	var result models.User
	err := db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, errors.New("invalid username or password")
		}
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return false, errors.New("invalid username or password")
	}

	return true, nil
}
