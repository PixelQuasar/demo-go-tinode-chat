package main

import (
	"demo-go-tinode-chat/internal/common/utils"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/routes"
	"fmt"
	"log"
	"os"
)

func main() {
	utils.LoadEnv()
	db.InitDB()

	server := routes.InitRouter()

	err := server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
}
