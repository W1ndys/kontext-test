package model

import (
	"time"
)

type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Slug      string    `gorm:"uniqueIndex;size:50;not null" json:"slug"`
	Articles  []Article `gorm:"many2many:article_tags;" json:"-"`
}
