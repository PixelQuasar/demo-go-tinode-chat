package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var UserCollection *mongo.Collection
var ChatConnection *mongo.Collection

func InitDB() {
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}

	database := client.Database("tinode_chat")

	UserCollection = database.Collection("users")
	ChatConnection = database.Collection("chat")
	
	fmt.Println("Connected to MongoDB!")
}
