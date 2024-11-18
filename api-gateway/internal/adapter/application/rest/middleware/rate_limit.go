package middleware

import (
	"api-gateway/internal/domain/port"

	"github.com/gin-gonic/gin"
)

func RateLimit(limiter port.ReteLimiterRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow(c, c.ClientIP()) {
			c.AbortWithStatusJSON(429, gin.H{"error": "too many request"})
			return
		}
		c.Next()
	}
}
