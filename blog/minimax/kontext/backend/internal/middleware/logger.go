package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Printf("[%s] %s %s %d %v %s",
			endTime.Format("2006-01-02 15:04:05"),
			method,
			path,
			statusCode,
			latency,
			clientIP,
		)
	}
}
