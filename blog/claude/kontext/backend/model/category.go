package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Articles []Article `gorm:"foreignKey:CategoryID" json:"articles,omitempty"`
}

type CategoryWithCount struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ArticleCount int64  `json:"article_count"`
}

func GetCategories(db *gorm.DB) ([]CategoryWithCount, error) {
	var categories []Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	var result []CategoryWithCount
	for _, cat := range categories {
		var count int64
		db.Model(&Article{}).Where("category_id = ? AND deleted_at IS NULL", cat.ID).Count(&count)
		result = append(result, CategoryWithCount{
			ID:           cat.ID,
			Name:         cat.Name,
			ArticleCount: count,
		})
	}

	return result, nil
}

func GetCategoryByID(db *gorm.DB, id uint) (*Category, error) {
	var category Category
	if err := db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(db *gorm.DB, category *Category) error {
	return db.Create(category).Error
}

func UpdateCategory(db *gorm.DB, id uint, name string) error {
	return db.Model(&Category{}).Where("id = ?", id).Update("name", name).Error
}

func DeleteCategory(db *gorm.DB, id uint) error {
	return db.Delete(&Category{}, id).Error
}

func CountArticlesByCategory(db *gorm.DB, categoryID uint) int64 {
	var count int64
	db.Model(&Article{}).Where("category_id = ? AND deleted_at IS NULL", categoryID).Count(&count)
	return count
}
