package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: implement authenticate
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "no authorization header"})
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		if !validateToken(token) {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}

func validateToken(token string) bool {
	// TODO: Implement JWT validation
	log.Println(token)
	return true
}
