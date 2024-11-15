package router

import "EStarExchange/api"

func (router RouterGroup) GoodGroupApi() {
	router.POST("/goodPicUpload", api.ApiGroupApp.GoodApi.GoodsPicUpload)
	router.POST("/goodInfoUpload", api.ApiGroupApp.GoodApi.GoodsInfoUpload)
	router.GET("/goodsGet", api.ApiGroupApp.GoodApi.GoodsGet)
	router.PUT("/goodsUpdate", api.ApiGroupApp.GoodApi.GoodsInfoUpdate)
	router.DELETE("/goodsDelete", api.ApiGroupApp.GoodApi.GoodsInfoDelete)
	router.POST("/goodsCollection", api.ApiGroupApp.GoodApi.GoodsCollection)
	router.GET("/addView", api.ApiGroupApp.GoodApi.GoodsAddView)
}
