package database

import (
	"log"

	"blog-backend/config"
	"blog-backend/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.AppConfig.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&model.Admin{}, &model.Category{}, &model.Tag{}, &model.Article{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	seedAdmin()
}

func seedAdmin() {
	var count int64
	DB.Model(&model.Admin{}).Count(&count)
	if count > 0 {
		return
	}

	hashedPassword, err := model.HashPassword("admin123")
	if err != nil {
		log.Fatalf("failed to hash default admin password: %v", err)
	}

	admin := model.Admin{
		Username: "admin",
		Password: hashedPassword,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatalf("failed to seed admin user: %v", err)
	}

	log.Println("default admin user created (username: admin, password: admin123)")
}
