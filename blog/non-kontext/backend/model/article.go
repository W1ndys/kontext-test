package model

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"type:text"`
	Category  string    `json:"category" gorm:"index"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Admin struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`
}
