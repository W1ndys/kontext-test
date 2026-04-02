package handler

import (
	"blog/internal/model/dto"
	"blog/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler() *CommentHandler {
	return &CommentHandler{
		commentService: service.NewCommentService(),
	}
}

func (h *CommentHandler) Create(c *gin.Context) {
	var req dto.CreateCommentDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	ip := c.ClientIP()
	comment, err := h.commentService.Create(&req, ip)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, comment)
}

func (h *CommentHandler) ListByArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的文章ID")
		return
	}

	comments, err := h.commentService.ListByArticleID(uint(id))
	if err != nil {
		InternalServerError(c, "获取评论列表失败")
		return
	}

	Success(c, comments)
}

func (h *CommentHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	status := c.Query("status")

	comments, total, err := h.commentService.List(page, pageSize, status)
	if err != nil {
		InternalServerError(c, "获取评论列表失败")
		return
	}

	Success(c, gin.H{
		"list":  comments,
		"total": total,
		"page":  page,
	})
}

func (h *CommentHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的评论ID")
		return
	}

	var req dto.UpdateCommentStatusDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误")
		return
	}

	comment, err := h.commentService.UpdateStatus(uint(id), req.Status)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, comment)
}

func (h *CommentHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequest(c, "无效的评论ID")
		return
	}

	if err := h.commentService.Delete(uint(id)); err != nil {
		InternalServerError(c, "删除评论失败")
		return
	}

	Success(c, nil)
}
