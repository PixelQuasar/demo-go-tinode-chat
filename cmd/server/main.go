package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type Book struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var collection *mongo.Collection

func loadEnvFiles() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}
}

func initMongo() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("tinode-chat").Collection("books")
}

func getBooks(c *gin.Context) {
	var books []Book
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching books"})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while adding a new book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book added", "book": book})
}

func main() {
	loadEnvFiles()
	initMongo()

	r := gin.Default()
	r.GET("/books", getBooks)
	r.POST("/books", addBook)

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
}
