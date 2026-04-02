package model

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ArticleID uint      `gorm:"not null;index" json:"article_id"`
	Article   Article   `gorm:"foreignKey:ArticleID" json:"-"`
	Nickname  string    `gorm:"size:50;not null" json:"nickname"`
	Email     string    `gorm:"size:100" json:"email"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	IP        string    `gorm:"size:50" json:"ip"`
	Status    string    `gorm:"size:20;default:'pending'" json:"status"` // pending, approved, rejected
}
