package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIPWhiteListMiddleware(t *testing.T) {
	// Initialize Gin
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Initialize IPWhitelist and Middleware
	ipWhitelist := NewIPWhiteList()
	ipWhitelist.AddIPs([]string{"127.0.0.1"}) // Allow only localhost

	r.Use(ipWhitelist.Middleware())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "allowed"})
	})

	// Test allowed IP
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "127.0.0.1:8080"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "allowed")

	// Test blocked IP
	req, _ = http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:8080"
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.Contains(t, w.Body.String(), "IP not allowed")
}
