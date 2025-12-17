package settings_api

import (
	"modulename/global"
	"modulename/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
	//c.JSONP(200, gin.H{"msg": "fdsfds"})
}
