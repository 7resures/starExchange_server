package router

import "EStarExchange/api"

func (router RouterGroup) UserInfoGroupApi() {
	router.GET("/userInfoGet", api.ApiGroupApp.UserApi.UserInfoGet)
	router.PUT("/updateUserInfo", api.ApiGroupApp.UserApi.UserInfoUpdate)
	router.PUT("/updatePwd", api.ApiGroupApp.UserApi.UserPwdUpdate)
	router.POST("/updateAvatar", api.ApiGroupApp.UserApi.AvatarUpdate)
}
