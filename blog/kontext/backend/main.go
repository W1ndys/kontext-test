package main

import (
	"fmt"
	"log"
	"os"

	"blog-backend/config"
	"blog-backend/database"
	"blog-backend/router"
)

func main() {
	config.LoadConfig()

	database.InitDB()

	if err := os.MkdirAll(config.AppConfig.UploadDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create upload directory: %v", err)
	}

	r := router.SetupRouter()

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	log.Printf("server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
