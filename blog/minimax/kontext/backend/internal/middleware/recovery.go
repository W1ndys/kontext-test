package middleware

import (
	"log"

	"blog/internal/handler"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v", r)
				handler.InternalServerError(c, "服务器内部错误")
				c.Abort()
			}
		}()
		c.Next()
	}
}
