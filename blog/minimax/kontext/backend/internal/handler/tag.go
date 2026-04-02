package handler

import (
	"blog/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	tagService *service.TagService
}

func NewTagHandler() *TagHandler {
	return &TagHandler{
		tagService: service.NewTagService(),
	}
}

func (h *TagHandler) Create(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
		Slug string `json:"slug"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误")
		return
	}

	tag, err := h.tagService.Create(req.Name, req.Slug)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, tag)
}

func (h *TagHandler) List(c *gin.Context) {
	tags, err := h.tagService.List()
	if err != nil {
		InternalServerError(c, "获取标签列表失败")
		return
	}

	Success(c, tags)
}

func (h *TagHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的标签ID")
		return
	}

	tag, err := h.tagService.GetByID(uint(id))
	if err != nil {
		NotFound(c, "标签不存在")
		return
	}

	Success(c, tag)
}

func (h *TagHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的标签ID")
		return
	}

	var req struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误")
		return
	}

	tag, err := h.tagService.Update(uint(id), req.Name, req.Slug)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, tag)
}

func (h *TagHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的标签ID")
		return
	}

	if err := h.tagService.Delete(uint(id)); err != nil {
		InternalServerError(c, "删除标签失败")
		return
	}

	Success(c, nil)
}
