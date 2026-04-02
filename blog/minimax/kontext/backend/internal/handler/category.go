package handler

import (
	"blog/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		categoryService: service.NewCategoryService(),
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Slug string `json:"slug"`
		Sort int    `json:"sort"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误")
		return
	}

	category, err := h.categoryService.Create(req.Name, req.Slug, req.Sort)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, category)
}

func (h *CategoryHandler) List(c *gin.Context) {
	categories, err := h.categoryService.List()
	if err != nil {
		InternalServerError(c, "获取分类列表失败")
		return
	}

	Success(c, categories)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的分类ID")
		return
	}

	category, err := h.categoryService.GetByID(uint(id))
	if err != nil {
		NotFound(c, "分类不存在")
		return
	}

	Success(c, category)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的分类ID")
		return
	}

	var req struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
		Sort int    `json:"sort"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误")
		return
	}

	category, err := h.categoryService.Update(uint(id), req.Name, req.Slug, req.Sort)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的分类ID")
		return
	}

	if err := h.categoryService.Delete(uint(id)); err != nil {
		InternalServerError(c, "删除分类失败")
		return
	}

	Success(c, nil)
}
