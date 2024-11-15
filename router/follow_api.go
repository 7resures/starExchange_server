package router

import "EStarExchange/api"

func (router RouterGroup) FollowGroupApi() {
	router.GET("/followGet", api.ApiGroupApp.FollowApi.FollowGet)
	router.PUT("/followUpdate", api.ApiGroupApp.FollowApi.FollowUpdate)
}
