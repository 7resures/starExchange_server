package api

import (
	"EStarExchange/api/advise_api"
	"EStarExchange/api/campus_api"
	"EStarExchange/api/image_api"
	"EStarExchange/api/login_api"
	"EStarExchange/api/user_api"
)

type ApiGroup struct {
	LoginApi  login_api.LoginApi
	ImageApi  image_api.ImageApi
	CampusApi campus_api.CampusApi
	UserApi   user_api.UserApi
	AdviseApi advise_api.AdviseApi
}

var ApiGroupApp = new(ApiGroup)
