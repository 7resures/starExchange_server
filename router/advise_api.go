package router

import "EStarExchange/api"

func (router RouterGroup) AdviseGroupApi() {
	router.POST("/updateAdvise", api.ApiGroupApp.AdviseApi.AdviseCreate)

}
