package api

import (
	"EStarExchange/api/campus_api"
	"EStarExchange/api/image_api"
	"EStarExchange/api/login_api"
)

type ApiGroup struct {
	LoginApi  login_api.LoginApi
	ImageApi  image_api.ImageApi
	CampusApi campus_api.CampusApi
}

var ApiGroupApp = new(ApiGroup)
