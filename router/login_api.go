package router

import "EStarExchange/api"

func (router RouterGroup) LoginGroupApi() {
	router.POST("/login", api.ApiGroupApp.LoginApi.LoginVerify)
	router.POST("/register", api.ApiGroupApp.LoginApi.RegisterHandler)
	router.GET("/campus_get", api.ApiGroupApp.CampusApi.CampusGet)
}
