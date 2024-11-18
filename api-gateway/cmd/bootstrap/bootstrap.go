package bootstrap

import (
	"api-gateway/internal/adapter/application/rest/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// api := "api"
	// version := "v1"

	router := gin.Default()

	// routerGroup := router.Group(path.Join(api, version))
	routerGroup := router.Group("/")

	// routes of the router
	routes.SetupRoutes(routerGroup)

	return router
}
