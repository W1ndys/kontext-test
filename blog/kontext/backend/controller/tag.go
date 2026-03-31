package controller

import (
	"net/http"
	"strconv"

	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

type CreateTagReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTagReq struct {
	Name string `json:"name" binding:"required"`
}

func GetTags(c *gin.Context) {
	tags, err := model.GetTags(database.DB)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to get tags")
		return
	}

	utils.Success(c, tags)
}

func CreateTag(c *gin.Context) {
	var req CreateTagReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	tag := model.Tag{Name: req.Name}
	if err := model.CreateTag(database.DB, &tag); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to create tag, name may already exist")
		return
	}

	utils.SuccessWithMessage(c, "tag created", tag)
}

func UpdateTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid tag id")
		return
	}

	if _, err := model.GetTagByID(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "tag not found")
		return
	}

	var req UpdateTagReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	if err := model.UpdateTag(database.DB, uint(id), req.Name); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to update tag")
		return
	}

	utils.SuccessWithMessage(c, "tag updated", gin.H{"id": id, "name": req.Name})
}

func DeleteTag(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid tag id")
		return
	}

	if _, err := model.GetTagByID(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "tag not found")
		return
	}

	if err := model.DeleteTag(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to delete tag")
		return
	}

	utils.SuccessWithMessage(c, "tag deleted", nil)
}
