package model

import (
	"time"
)

type Article struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Title      string    `gorm:"size:200;not null" json:"title"`
	Slug       string    `gorm:"uniqueIndex;size:200" json:"slug"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	Summary    string    `gorm:"size:500" json:"summary"`
	CoverImage string    `gorm:"size:500" json:"cover_image"`
	Status     string    `gorm:"size:20;default:'draft'" json:"status"` // draft, published
	ViewCount  int       `gorm:"default:0" json:"view_count"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Tags       []Tag     `gorm:"many2many:article_tags;" json:"tags,omitempty"`
}
