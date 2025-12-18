package routers

import (
	"modulename/api"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("login", userApi.EmailLoginView)
	router.POST("Register", userApi.Register)
}
