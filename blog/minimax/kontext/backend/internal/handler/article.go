package handler

import (
	"blog/internal/model/dto"
	"blog/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *service.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		articleService: service.NewArticleService(),
	}
}

func (h *ArticleHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	categoryID, _ := strconv.ParseUint(c.Query("category_id"), 10, 32)
	tagID, _ := strconv.ParseUint(c.Query("tag_id"), 10, 32)
	status := c.DefaultQuery("status", "published")

	articles, total, err := h.articleService.List(page, pageSize, uint(categoryID), uint(tagID), status)
	if err != nil {
		InternalServerError(c, "获取文章列表失败")
		return
	}

	Success(c, gin.H{
		"list":  articles,
		"total": total,
		"page":  page,
	})
}

func (h *ArticleHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的文章ID")
		return
	}

	article, err := h.articleService.GetByID(uint(id))
	if err != nil {
		NotFound(c, "文章不存在")
		return
	}

	Success(c, article)
}

func (h *ArticleHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	article, err := h.articleService.GetBySlug(slug)
	if err != nil {
		NotFound(c, "文章不存在")
		return
	}

	Success(c, article)
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req dto.CreateArticleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	article, err := h.articleService.Create(&req)
	if err != nil {
		InternalServerError(c, "创建文章失败")
		return
	}

	Success(c, article)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的文章ID")
		return
	}

	var req dto.UpdateArticleDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	article, err := h.articleService.Update(uint(id), &req)
	if err != nil {
		InternalServerError(c, "更新文章失败")
		return
	}

	Success(c, article)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的文章ID")
		return
	}

	if err := h.articleService.Delete(uint(id)); err != nil {
		InternalServerError(c, "删除文章失败")
		return
	}

	Success(c, nil)
}

func (h *ArticleHandler) ListAll(c *gin.Context) {
	articles, err := h.articleService.ListAll()
	if err != nil {
		InternalServerError(c, "获取文章列表失败")
		return
	}

	Success(c, articles)
}
