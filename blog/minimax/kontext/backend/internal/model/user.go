package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Nickname  string    `gorm:"size:100" json:"nickname"`
	Avatar    string    `gorm:"size:500" json:"avatar"`
}
