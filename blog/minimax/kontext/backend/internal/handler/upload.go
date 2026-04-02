package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"blog/internal/config"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	config *config.Config
}

func NewUploadHandler() *UploadHandler {
	cfg := config.GlobalConfig
	if cfg == nil {
		cfg = &config.Config{UploadPath: "./uploads", MaxImageSize: 5242880}
	}
	return &UploadHandler{config: cfg}
}

func (h *UploadHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		BadRequest(c, "未上传文件")
		return
	}

	if file.Size > h.config.MaxImageSize {
		BadRequest(c, fmt.Sprintf("文件大小不能超过 %d MB", h.config.MaxImageSize/1024/1024))
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" && ext != ".webp" {
		BadRequest(c, "仅支持 JPG、PNG、GIF、WebP 格式图片")
		return
	}

	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), randomString(8), ext)
	filepath := filepath.Join(h.config.UploadPath, filename)

	if err := os.MkdirAll(h.config.UploadPath, 0755); err != nil {
		InternalServerError(c, "创建上传目录失败")
		return
	}

	if err := c.SaveUploadedFile(file, filepath); err != nil {
		InternalServerError(c, "保存文件失败")
		return
	}

	Success(c, gin.H{
		"url":      "/uploads/" + filename,
		"filename": filename,
	})
}

func (h *UploadHandler) ServeStatic(c *gin.Context) {
	c.File(filepath.Join(h.config.UploadPath, c.Param("filename")))
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(result)
}
