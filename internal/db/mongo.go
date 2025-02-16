package db

import (
	"context"
	"demo-go-tinode-chat/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var ChatConnection *mongo.Collection
var TinodeAuthCollection *mongo.Collection

func InitDB() {
	clientOpts := options.Client().ApplyURI(config.AppConfig.MongoUri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}

	database := client.Database("tinode_chat")

	UserCollection = database.Collection("users")
	ChatConnection = database.Collection("chat")
	TinodeAuthCollection = database.Collection("auth")

	fmt.Println("Connected to MongoDB!")
}
