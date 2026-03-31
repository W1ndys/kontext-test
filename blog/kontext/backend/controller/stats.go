package controller

import (
	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

func GetStats(c *gin.Context) {
	var articleCount int64
	database.DB.Model(&model.Article{}).Count(&articleCount)

	var categoryCount int64
	database.DB.Model(&model.Category{}).Count(&categoryCount)

	var tagCount int64
	database.DB.Model(&model.Tag{}).Count(&tagCount)

	totalViews := model.GetTotalViewCount(database.DB)

	utils.Success(c, gin.H{
		"article_count":  articleCount,
		"category_count": categoryCount,
		"tag_count":      tagCount,
		"total_views":    totalViews,
	})
}
