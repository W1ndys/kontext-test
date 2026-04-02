package controller

import (
	"net/http"

	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	admin, err := model.GetAdminByUsername(database.DB, req.Username)
	if err != nil {
		utils.Fail(c, http.StatusUnauthorized, 401, "invalid username or password")
		return
	}

	if !admin.CheckPassword(req.Password) {
		utils.Fail(c, http.StatusUnauthorized, 401, "invalid username or password")
		return
	}

	token, err := utils.GenerateToken(admin.ID, admin.Username)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to generate token")
		return
	}

	utils.Success(c, gin.H{
		"token":    token,
		"username": admin.Username,
	})
}

func GetProfile(c *gin.Context) {
	adminID, _ := c.Get("admin_id")
	username, _ := c.Get("username")

	utils.Success(c, gin.H{
		"id":       adminID,
		"username": username,
	})
}
