package handler

import (
	"api-gateway/internal/domain/service"
	"io"
	"net/http"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read request body"})
		return
	}

	// fmt.Println("body: ", body) // in bytes
	// fmt.Println("method: ", c.Request.Method)

	response, err := h.gatewayService.ProxyRequest(
		c.Request.Context(),
		c.Request.URL.Path,
		c.Request.Method,
		c.Request.Header,
		body,
	)

	// fmt.Println("response: ", response)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set headers from the response
	for key, values := range response.Headers {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	// Forward the status code and body to the client
	c.Data(response.StatusCode, c.ContentType(), response.Body)
}
