package router

import "EStarExchange/api"

func (router RouterGroup) TagGroupApi() {
	router.POST("/tagCreate", api.ApiGroupApp.TagApi.TagCreate)

}
