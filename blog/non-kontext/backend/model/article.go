package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null"`
	Content  string `json:"content" gorm:"type:text"`
	Category string `json:"category" gorm:"index"`
	Tags     string `json:"tags"`     // comma-separated
	Summary  string `json:"summary"`  // short excerpt
	Cover    string `json:"cover"`    // cover image URL
	Status   string `json:"status" gorm:"default:published"` // published / draft
}

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"` // bcrypt hash, never serialize
}
