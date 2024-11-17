package routes

import (
	"api-gateway/internal/adapter/application/rest/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(proxyHandler *handler.ProxyHandler, rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")

	userGroup.Any("/*path", proxyHandler.Handle)
}
