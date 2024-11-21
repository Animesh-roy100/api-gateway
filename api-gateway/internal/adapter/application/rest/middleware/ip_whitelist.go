package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type IPWhiteList struct {
	whitelist map[string]bool
	mu        sync.RWMutex
}

func NewIPWhiteList() *IPWhiteList {
	return &IPWhiteList{
		whitelist: make(map[string]bool),
	}
}

// working as a middleware
func (w *IPWhiteList) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !w.IsIPWhitelisted(clientIP) {
			c.JSON(http.StatusForbidden, gin.H{"error": "IP not allowed", "ip": clientIP})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (w *IPWhiteList) AddIPs(ips []string) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if len(ips) > 0 && ips[0] == "ALL" {
		// Allow all IP ranges
		w.whitelist["ALL"] = true
		return
	}

	newList := make(map[string]bool)
	for _, ip := range ips {
		if ip == "ALL" {
			continue
		}
		newList[ip] = true
	}

	w.whitelist = newList
}

func (w *IPWhiteList) GetWhiteList() map[string]bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	// make a copy to prevent external modification
	copied := make(map[string]bool)
	for ip, allowed := range w.whitelist {
		copied[ip] = allowed
	}

	return w.whitelist
}

func (w *IPWhiteList) UpdateWhiteList(newList map[string]bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.whitelist = make(map[string]bool)

	for ip, allowed := range newList {
		w.whitelist[ip] = allowed
	}
}

func (w *IPWhiteList) IsIPWhitelisted(ip string) bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if _, exists := w.whitelist["ALL"]; exists {
		return true
	}

	return w.whitelist[ip]
}

func (w *IPWhiteList) AddIP(ip string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.whitelist[ip] = true
}

func (w *IPWhiteList) RemoveIP(ip string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	delete(w.whitelist, ip)
}
