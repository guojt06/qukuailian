package api

import (
	"modulename/api/settings_api"
	"modulename/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	UserApi     user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
