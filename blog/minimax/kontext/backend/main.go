package main

import (
	"blog/internal/config"
	"blog/internal/repository"
	"blog/internal/router"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg := config.LoadConfig()

	repository.GetDB()

	// Initialize default admin user with proper bcrypt hash
	if err := initDefaultAdmin(); err != nil {
		log.Printf("Warning: Failed to initialize admin user: %v", err)
	}

	if err := repository.InitDefaultData(); err != nil {
		log.Printf("Warning: Failed to initialize default data: %v", err)
	}

	if cfg.ServerPort == "" {
		cfg.ServerPort = ":8080"
	}

	r := router.SetupRouter()

	fmt.Printf("Server starting on %s\n", cfg.ServerPort)
	if err := r.Run(cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initDefaultAdmin() error {
	db := repository.GetDB()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		_, err = db.Exec(
			"INSERT INTO users (username, password, nickname) VALUES (?, ?, ?)",
			"admin", string(hash), "博主",
		)
		if err != nil {
			return err
		}
		fmt.Println("Admin user created: admin / admin123")
	}
	return nil
}
