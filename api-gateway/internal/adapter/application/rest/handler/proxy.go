package handler

import (
	"api-gateway/internal/domain/service"
	"io"

	"github.com/gin-gonic/gin"
)

type ProxyHandler struct {
	gatewayService service.GatewayService
}

func NewProxyHandler(gs service.GatewayService) *ProxyHandler {
	return &ProxyHandler{
		gatewayService: gs,
	}
}

func (h *ProxyHandler) Handle(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read request body"})
		return
	}

	response, err := h.gatewayService.ProxyRequest(
		c.Request.Context(),
		c.Request.URL.Path,
		c.Request.Method,
		c.Request.Header,
		body,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for key, values := range response.Headers {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Data(response.StatusCode, "application/json", response.Body)
}
