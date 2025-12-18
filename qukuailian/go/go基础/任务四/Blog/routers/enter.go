package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()

	apiRouters := routers.Group(`api`)
	routerGroupApp := RouterGroup{apiRouters}

	routerGroupApp.SettingsRouter()
	routerGroupApp.UserRouter()
	return routers
}
