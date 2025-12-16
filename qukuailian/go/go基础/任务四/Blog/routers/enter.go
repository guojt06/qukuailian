package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routers := gin.Default()
	routers.GET("/", func(c *gin.Context) {
		c.JSON(200, "fdsfdsfs ")
	})
	return routers
}
