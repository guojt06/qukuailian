package settings_api

import (
	"modulename/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OK(map[string]string{}, "dsdaa", c)
	//c.JSONP(200, gin.H{"msg": "fdsfds"})
}
