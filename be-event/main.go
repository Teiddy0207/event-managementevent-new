package main

import (
	"be-event/config"
	"be-event/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//	log.Fatal("❌ Error loading .env file")
	//}

	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("⚠️ Không tìm thấy .env.local, đang dùng .env mặc định")
		if err := godotenv.Load(); err != nil {
			log.Fatal("❌ Không thể load .env")
		}
	}
	config.ConnectDatabase()

	r := routes.InitRouter()

	r.Run() // chạy ở PORT env hoặc mặc định 8080
}
