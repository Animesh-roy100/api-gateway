package routes

import (
	"api-gateway/internal/adapter/application/rest/handler"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(proxyHandler *handler.ProxyHandler, rg *gin.RouterGroup) {
	productGroup := rg.Group("/products")

	productGroup.Any("/*path", proxyHandler.Handle)
}
