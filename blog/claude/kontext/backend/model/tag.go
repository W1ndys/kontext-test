package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"name"`
	Articles []Article `gorm:"many2many:article_tags" json:"articles,omitempty"`
}

type TagWithCount struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	ArticleCount int64  `json:"article_count"`
}

func GetTags(db *gorm.DB) ([]TagWithCount, error) {
	var tags []Tag
	if err := db.Find(&tags).Error; err != nil {
		return nil, err
	}

	var result []TagWithCount
	for _, tag := range tags {
		var count int64
		db.Table("article_tags").
			Joins("JOIN articles ON articles.id = article_tags.article_id").
			Where("article_tags.tag_id = ? AND articles.deleted_at IS NULL", tag.ID).
			Count(&count)
		result = append(result, TagWithCount{
			ID:           tag.ID,
			Name:         tag.Name,
			ArticleCount: count,
		})
	}

	return result, nil
}

func GetTagByID(db *gorm.DB, id uint) (*Tag, error) {
	var tag Tag
	if err := db.First(&tag, id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func CreateTag(db *gorm.DB, tag *Tag) error {
	return db.Create(tag).Error
}

func UpdateTag(db *gorm.DB, id uint, name string) error {
	return db.Model(&Tag{}).Where("id = ?", id).Update("name", name).Error
}

func DeleteTag(db *gorm.DB, id uint) error {
	// Remove associations first
	db.Exec("DELETE FROM article_tags WHERE tag_id = ?", id)
	return db.Delete(&Tag{}, id).Error
}
