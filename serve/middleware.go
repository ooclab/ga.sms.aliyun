package serve

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func gwAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 执行任何检查
		c.Next()
	}
}
