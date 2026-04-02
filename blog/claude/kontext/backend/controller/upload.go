package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"blog-backend/config"
	"blog-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var allowedImageTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

const maxImageSize = 5 << 20 // 5MB

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, 400, "no image file provided")
		return
	}

	if file.Size > maxImageSize {
		utils.Fail(c, http.StatusBadRequest, 400, "image size exceeds 5MB limit")
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageTypes[ext] {
		utils.Fail(c, http.StatusBadRequest, 400, "unsupported image type, allowed: jpg, jpeg, png, gif, webp")
		return
	}

	if err := os.MkdirAll(config.AppConfig.UploadDir, os.ModePerm); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to create upload directory")
		return
	}

	filename := uuid.New().String() + ext
	savePath := filepath.Join(config.AppConfig.UploadDir, filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		utils.Fail(c, http.StatusInternalServerError, 500, "failed to save image")
		return
	}

	url := fmt.Sprintf("/uploads/%s", filename)

	utils.Success(c, gin.H{
		"url":      url,
		"filename": filename,
	})
}
