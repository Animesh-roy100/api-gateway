package bootstrap

import (
	"api-gateway/internal/adapter/application/rest/routes"
	"api-gateway/internal/middleware"
	"path"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	api := "api"
	version := "v1"

	router := gin.Default()

	// Apply middlewares
	router.Use(middleware.RateLimit())
	router.Use(middleware.Authenticate())

	routerGroup := router.Group(path.Join(api, version))

	// routes of the router
	routes.SetupRoutes(routerGroup)

	return router
}
