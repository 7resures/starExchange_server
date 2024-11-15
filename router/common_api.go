package router

import "EStarExchange/api"

func (router RouterGroup) commonGroupApi() {
	router.GET("/tagGet", api.ApiGroupApp.TagApi.TagGet)
}
