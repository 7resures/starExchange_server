package router

import "EStarExchange/api"

func (router RouterGroup) ImagesUploadGroupApi() {
	router.POST("/imagesUpload", api.ApiGroupApp.ImageApi.ImagesUplpoad)
}
