package settings_api

import (
	"modulename/config"
	"modulename/core"
	"modulename/global"
	"modulename/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsUpda(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("参数错误", c)
		return
	}

	global.Log.Info("server运行在: ", global.Config)
	//fmt.Sprint("修改之前的数据", global.Config)
	global.Config.SiteInfo = cr
	//fmt.Sprint("修改之后的数据", global.Config)
	global.Log.Info("server运行在: ", global.Config)
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OKWithnil(c)

}
