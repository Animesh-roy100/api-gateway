package bootstrap

import (
	"path"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	api := "api"
	version := "v1"

	router := gin.New()

	router.Group(path.Join(api, version))

	return router
}
