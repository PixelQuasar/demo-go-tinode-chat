package main

import (
	"demo-go-tinode-chat/config"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/server"
	"demo-go-tinode-chat/internal/tinode"
	"fmt"
	"log"
)

func main() {
	config.LoadConfig()
	db.InitDB()
	tinode.InitMessageLoop()

	app := server.InitRouter()

	err := app.Run(fmt.Sprintf(":%s", config.AppConfig.Port))
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
}
