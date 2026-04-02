package model

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Slug      string    `gorm:"uniqueIndex;size:50;not null" json:"slug"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Articles  []Article `gorm:"foreignKey:CategoryID" json:"-"`
}
