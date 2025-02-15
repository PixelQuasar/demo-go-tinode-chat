package handlers

import (
	"demo-go-tinode-chat/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func PostMessageHandler(c *gin.Context) {
	fmt.Println("hello there")
	var requestData struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	fmt.Println("hello there", requestData.Content)

	authHeader := c.GetHeader("Authorization")
	fmt.Println("hello there", authHeader)

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
		return
	}
	userToken := strings.TrimPrefix(authHeader, "Bearer ")

	fmt.Println("hello there", userToken)

	err := services.CreateMessage(requestData.Content, userToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message created successfully"})
}

func MessagesHandler(c *gin.Context) {
	messages, err := services.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
