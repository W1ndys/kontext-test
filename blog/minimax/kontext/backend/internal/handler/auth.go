package handler

import (
	"blog/internal/model/dto"
	"blog/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		BadRequest(c, err.Error())
		return
	}

	Success(c, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequest(c, "请求参数错误: "+err.Error())
		return
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		Unauthorized(c, err.Error())
		return
	}

	Success(c, resp)
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		Unauthorized(c, "未登录")
		return
	}

	user, err := h.authService.GetUserByID(userID.(uint))
	if err != nil {
		Unauthorized(c, "获取用户信息失败")
		return
	}

	Success(c, user)
}
