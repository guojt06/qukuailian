package routers

import (
	"modulename/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("setttings", settingsApi.SettingsInfoView)
}
