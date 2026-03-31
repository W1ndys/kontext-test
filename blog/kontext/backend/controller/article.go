package controller

import (
	"net/http"
	"strconv"

	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

type CreateArticleReq struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Summary    string `json:"summary"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids"`
	Published  *bool  `json:"published"`
}

type UpdateArticleReq struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Summary    string `json:"summary"`
	CategoryID uint   `json:"category_id"`
	TagIDs     []uint `json:"tag_ids"`
	Published  *bool  `json:"published"`
}

func GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryID, _ := strconv.ParseUint(c.Query("category_id"), 10, 32)
	tagID, _ := strconv.ParseUint(c.Query("tag_id"), 10, 32)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	articles, total := model.GetArticles(database.DB, page, pageSize, uint(categoryID), uint(tagID))

	utils.Success(c, gin.H{
		"list":      articles,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func GetArticleDetail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid article id")
		return
	}

	article, err := model.GetArticleByID(database.DB, uint(id))
	if err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "article not found")
		return
	}

	_ = model.IncrementViewCount(database.DB, uint(id))
	article.ViewCount++

	utils.Success(c, article)
}

func CreateArticle(c *gin.Context) {
	var req CreateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	published := true
	if req.Published != nil {
		published = *req.Published
	}

	article := model.Article{
		Title:      req.Title,
		Content:    req.Content,
		Summary:    req.Summary,
		CategoryID: req.CategoryID,
		Published:  published,
	}

	if err := model.CreateArticle(database.DB, &article, req.TagIDs); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to create article")
		return
	}

	// Reload with associations
	created, _ := model.GetArticleByID(database.DB, article.ID)
	utils.SuccessWithMessage(c, "article created", created)
}

func UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid article id")
		return
	}

	article, err := model.GetArticleByID(database.DB, uint(id))
	if err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "article not found")
		return
	}

	var req UpdateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	if req.Title != "" {
		article.Title = req.Title
	}
	if req.Content != "" {
		article.Content = req.Content
	}
	if req.Summary != "" {
		article.Summary = req.Summary
	}
	if req.CategoryID > 0 {
		article.CategoryID = req.CategoryID
	}
	if req.Published != nil {
		article.Published = *req.Published
	}

	if err := model.UpdateArticle(database.DB, article, req.TagIDs); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to update article")
		return
	}

	updated, _ := model.GetArticleByID(database.DB, article.ID)
	utils.SuccessWithMessage(c, "article updated", updated)
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid article id")
		return
	}

	if _, err := model.GetArticleByID(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "article not found")
		return
	}

	if err := model.DeleteArticle(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to delete article")
		return
	}

	utils.SuccessWithMessage(c, "article deleted", nil)
}

func SearchArticles(c *gin.Context) {
	keyword := c.Query("keyword")
	if keyword == "" {
		utils.Fail(c, http.StatusBadRequest, 400, "keyword is required")
		return
	}

	articles, err := model.SearchArticles(database.DB, keyword)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "search failed")
		return
	}

	utils.Success(c, articles)
}

func GetArticlesTimeline(c *gin.Context) {
	timeline, err := model.GetArticlesTimeline(database.DB)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to get timeline")
		return
	}

	utils.Success(c, timeline)
}
