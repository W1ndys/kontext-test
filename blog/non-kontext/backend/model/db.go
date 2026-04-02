package model

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Ensure data directory exists
	if err := os.MkdirAll("data", 0755); err != nil {
		log.Fatalf("failed to create data directory: %v", err)
	}

	var err error
	DB, err = gorm.Open(sqlite.Open("data/blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto-migrate schemas
	if err := DB.AutoMigrate(&Article{}, &Admin{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Seed default admin if none exists
	var count int64
	DB.Model(&Admin{}).Count(&count)
	if count == 0 {
		hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("failed to hash default password: %v", err)
		}
		admin := Admin{
			Username: "admin",
			Password: string(hash),
		}
		if err := DB.Create(&admin).Error; err != nil {
			log.Fatalf("failed to create default admin: %v", err)
		}
		log.Println("Default admin created (admin / admin123)")
	}
}
