package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title      string   `gorm:"type:varchar(200);not null" json:"title"`
	Content    string   `gorm:"type:text;not null" json:"content"`
	Summary    string   `gorm:"type:varchar(500)" json:"summary"`
	CategoryID uint     `gorm:"index;not null" json:"category_id"`
	ViewCount  int      `gorm:"default:0" json:"view_count"`
	Published  bool     `gorm:"default:true" json:"published"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
	Tags       []Tag    `gorm:"many2many:article_tags" json:"tags"`
}

type ArticleTimeline struct {
	Month    string    `json:"month"`
	Articles []Article `json:"articles"`
}

func GetArticles(db *gorm.DB, page, pageSize int, categoryID, tagID uint) ([]Article, int64) {
	var articles []Article
	var total int64

	query := db.Model(&Article{})

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if tagID > 0 {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Where("article_tags.tag_id = ?", tagID)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	query.Preload("Category").Preload("Tags").
		Order("created_at DESC").
		Offset(offset).Limit(pageSize).
		Find(&articles)

	return articles, total
}

func GetArticleByID(db *gorm.DB, id uint) (*Article, error) {
	var article Article
	if err := db.Preload("Category").Preload("Tags").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func CreateArticle(db *gorm.DB, article *Article, tagIDs []uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(article).Error; err != nil {
			return err
		}

		if len(tagIDs) > 0 {
			var tags []Tag
			if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
				return err
			}
			if err := tx.Model(article).Association("Tags").Replace(tags); err != nil {
				return err
			}
		}

		return nil
	})
}

func UpdateArticle(db *gorm.DB, article *Article, tagIDs []uint) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(article).Error; err != nil {
			return err
		}

		var tags []Tag
		if len(tagIDs) > 0 {
			if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(article).Association("Tags").Replace(tags); err != nil {
			return err
		}

		return nil
	})
}

func DeleteArticle(db *gorm.DB, id uint) error {
	// Remove tag associations first
	db.Exec("DELETE FROM article_tags WHERE article_id = ?", id)
	return db.Delete(&Article{}, id).Error
}

func SearchArticles(db *gorm.DB, keyword string) ([]Article, error) {
	var articles []Article
	likePattern := "%" + keyword + "%"
	if err := db.Preload("Category").Preload("Tags").
		Where("title LIKE ? OR content LIKE ?", likePattern, likePattern).
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticlesByTag(db *gorm.DB, tagID uint) ([]Article, error) {
	var articles []Article
	if err := db.Preload("Category").Preload("Tags").
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID).
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticlesTimeline(db *gorm.DB) ([]ArticleTimeline, error) {
	var articles []Article
	if err := db.Preload("Category").Preload("Tags").
		Order("created_at DESC").
		Find(&articles).Error; err != nil {
		return nil, err
	}

	timelineMap := make(map[string][]Article)
	var order []string

	for _, article := range articles {
		month := fmt.Sprintf("%d-%02d", article.CreatedAt.Year(), article.CreatedAt.Month())
		if _, exists := timelineMap[month]; !exists {
			order = append(order, month)
		}
		timelineMap[month] = append(timelineMap[month], article)
	}

	var timeline []ArticleTimeline
	for _, month := range order {
		timeline = append(timeline, ArticleTimeline{
			Month:    month,
			Articles: timelineMap[month],
		})
	}

	return timeline, nil
}

func IncrementViewCount(db *gorm.DB, id uint) error {
	return db.Model(&Article{}).Where("id = ?", id).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error
}

func GetTotalViewCount(db *gorm.DB) int64 {
	var total int64
	row := db.Model(&Article{}).Select("COALESCE(SUM(view_count), 0)").Row()
	row.Scan(&total)
	return total
}
