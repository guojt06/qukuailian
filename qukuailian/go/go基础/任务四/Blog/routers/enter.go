package routers

import (
	"modulename/utils/auth"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()

	apiRouters := routers.Group(`api`)
	//apiRouters.Use(auth.AdminAuthMiddleware())
	routerGroupApp := RouterGroup{apiRouters}
	routerGroupApp.SettingsRouter()
	routerGroupApp.UserRouter()

	//global.BackendRouter = global.Engine.Group("/backend")
	//global.BackendRouter.Use(auth.AdminAuthMiddleware())

	post := routers.Group(`post`)
	post.Use(auth.AdminAuthMiddleware())
	postApp := RouterGroup{post}
	postApp.RouterPostApp()

	return routers
}
