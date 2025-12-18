package post_api

import (
	"modulename/global"
	"modulename/models/res"

	"github.com/gin-gonic/gin"
)

func (PostApi) PostInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
