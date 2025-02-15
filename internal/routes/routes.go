package routes

import (
	"demo-go-tinode-chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", handlers.SignupHandler)
	router.POST("/login", handlers.LoginHandler)

	return router
}
