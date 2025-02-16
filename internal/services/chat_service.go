package services

import (
	"context"
	"demo-go-tinode-chat/config"
	"demo-go-tinode-chat/internal/common/utils"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/models"
	"demo-go-tinode-chat/internal/tinode"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func CreateMessage(content string, userToken string) error {
	username, err := utils.GetUsernameByToken(userToken)
	if err != nil {
		return err
	}

	fmt.Println("message creating")

	message := models.Message{
		Content:   content,
		Author:    username,
		Timestamp: time.Now(),
	}

	_, err = db.ChatConnection.InsertOne(context.TODO(), message)
	if err != nil {
		return err
	}

	err = tinode.SendMessage(content, username)
	return err
}

func GetMessages() ([]models.Message, error) {
	var messages []models.Message

	findOptions := options.Find()
	findOptions.SetLimit(config.AppConfig.MessagesPageSize)
	findOptions.SetSort(bson.D{{"timestamp", 1}})

	cursor, err := db.ChatConnection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, context.Background())

	if err = cursor.All(context.Background(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
