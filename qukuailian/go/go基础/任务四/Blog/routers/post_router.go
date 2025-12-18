package routers

import (
	"modulename/api"
)

func (router RouterGroup) RouterPostApp() {
	psotApi := api.ApiGroupApp.PostApi
	router.POST("/users", psotApi.PostInfoView)

}
