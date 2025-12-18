package api

import (
	"modulename/api/post_api"
	"modulename/api/settings_api"
	"modulename/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	UserApi     user_api.UserApi
	PostApi     post_api.PostApi
}

var ApiGroupApp = new(ApiGroup)
