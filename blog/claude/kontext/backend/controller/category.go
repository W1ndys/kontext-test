package controller

import (
	"net/http"
	"strconv"

	"blog-backend/database"
	"blog-backend/model"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

type CreateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryReq struct {
	Name string `json:"name" binding:"required"`
}

func GetCategories(c *gin.Context) {
	categories, err := model.GetCategories(database.DB)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to get categories")
		return
	}

	utils.Success(c, categories)
}

func CreateCategory(c *gin.Context) {
	var req CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	category := model.Category{Name: req.Name}
	if err := model.CreateCategory(database.DB, &category); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to create category, name may already exist")
		return
	}

	utils.SuccessWithMessage(c, "category created", category)
}

func UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid category id")
		return
	}

	if _, err := model.GetCategoryByID(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "category not found")
		return
	}

	var req UpdateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid request parameters")
		return
	}

	if err := model.UpdateCategory(database.DB, uint(id), req.Name); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to update category")
		return
	}

	utils.SuccessWithMessage(c, "category updated", gin.H{"id": id, "name": req.Name})
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "invalid category id")
		return
	}

	if _, err := model.GetCategoryByID(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusNotFound, 404, "category not found")
		return
	}

	count := model.CountArticlesByCategory(database.DB, uint(id))
	if count > 0 {
		utils.Fail(c, http.StatusBadRequest, 400, "cannot delete category with existing articles")
		return
	}

	if err := model.DeleteCategory(database.DB, uint(id)); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to delete category")
		return
	}

	utils.SuccessWithMessage(c, "category deleted", nil)
}
