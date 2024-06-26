package middlewares

import (
	"go-web-app/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 示例权限检查函数
func checkPermission(c *gin.Context) bool {
	token := c.GetHeader("Authorization")
	if token == "" {
		return false
	}
	// 假设 token 应该是 "valid-token"
	return token == "valid-token"
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkPermission(c) {
			logger.Logger.Warn("Unauthorized access")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
