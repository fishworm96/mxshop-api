package initialize

import (
	"mxshop-api/user-web/middlewares"
	router "mxshop-api/user-web/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())

	zap.S().Info("配置用户相关的url")
	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}