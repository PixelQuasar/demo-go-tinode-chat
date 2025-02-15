package handlers

import (
	auth "demo-go-tinode-chat/internal/common/utils"
	"demo-go-tinode-chat/internal/models"
	"demo-go-tinode-chat/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func LoginHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result bool
	result, err := services.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error authenticating user"})
		return
	}

	if !result {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	var token string
	token, err = auth.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
