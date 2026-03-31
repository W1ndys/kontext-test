package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(200);not null" json:"-"`
}

func GetAdminByUsername(db *gorm.DB, username string) (*Admin, error) {
	var admin Admin
	if err := db.Where("username = ?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
