package server

import (
	auth "demo-go-tinode-chat/internal/common/middleware"
	"demo-go-tinode-chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", handlers.SignupHandler)
	router.POST("/login", handlers.LoginHandler)

	authorized := router.Group("/")
	authorized.Use(auth.Middleware())
	{
		authorized.GET("/messages", handlers.MessagesHandler)
		authorized.POST("/message", handlers.PostMessageHandler)
	}

	return router
}
