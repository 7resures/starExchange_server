package router

import (
	"EStarExchange/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//安装gin路由框架
//go get github.com/gin-gonic/gin

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//图片访问api
	router.Static("/avatars", "./uploadFile/avatar")
	router.Static("/images", "./uploadFile/image")
	//前台相关API
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.LoginGroupApi()
	routerGroupApp.ImagesUploadGroupApi()
	routerGroupApp.UserInfoGroupApi()
	routerGroupApp.AdviseGroupApi()
	routerGroupApp.GoodGroupApi()
	routerGroupApp.FollowGroupApi()

	// 公共Api
	commonRouterGroup := router.Group("common")
	commonGroupApp := RouterGroup{commonRouterGroup}
	commonGroupApp.commonGroupApi()

	//后台管理相关API
	adminRouterGroup := router.Group("admin")
	adminGroupApp := RouterGroup{adminRouterGroup}
	adminGroupApp.CampusGroupApi()
	adminGroupApp.TagGroupApi()

	return router
}
