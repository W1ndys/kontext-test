package middleware

import (
	"net/http"
	"strings"

	"blog-backend/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Fail(c, http.StatusUnauthorized, 401, "missing authorization header")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Fail(c, http.StatusUnauthorized, 401, "invalid authorization header format")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			utils.Fail(c, http.StatusUnauthorized, 401, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set("admin_id", claims.AdminID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
