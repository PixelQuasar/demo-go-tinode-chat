package handlers

import (
	"demo-go-tinode-chat/internal/common/utils"
	"demo-go-tinode-chat/internal/models"
	"demo-go-tinode-chat/internal/services"
	"demo-go-tinode-chat/internal/tinode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignupHandler(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := utils.ValidatePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password: it must be at least 8 characters " +
				"long and contain at least one letter and one digit",
		})
		return
	}

	exists, err := utils.CheckIfUserExists(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking if user exists"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	err = services.CreateUser(user)
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

	result, err := services.AuthenticateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error authenticating user"})
		return
	}

	if !result {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	err = tinode.CreateTinodeUser(user, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating Tinode user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
