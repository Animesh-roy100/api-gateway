package routes

import (
	"api-gateway/internal/adapter/application/rest/handler"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(proxyHandler *handler.ProxyHandler, rg *gin.RouterGroup) {
	paymentGroup := rg.Group("/payments")

	paymentGroup.Any("/*path", proxyHandler.Handle)
}
