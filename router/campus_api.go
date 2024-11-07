package router

import (
	"EStarExchange/api"
)

func (router RouterGroup) CampusGroupApi() {
	router.POST("/campusCreate", api.ApiGroupApp.CampusApi.CreateCampus)
}
