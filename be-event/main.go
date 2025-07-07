package main

import (
	"be-event/config"
	"be-event/routes"

	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	config.ConnectDatabase()

	r := routes.InitRouter()

	r.Run() // chạy ở PORT env hoặc mặc định 8080
}
