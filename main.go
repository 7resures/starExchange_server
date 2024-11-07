package main

import (
	"EStarExchange/core"
	_ "EStarExchange/docs"
	"EStarExchange/flag"
	"EStarExchange/global"
	"EStarExchange/router"
)

// @title estarexchange API文档
// @version V1.0
// @description estarexchange的API文档
// @host 127.0.0.1:8080
func main() {
	//初始化配置文件
	core.InitConf()

	//初始化日志
	global.Log = core.InitLogger()
	//初始化gorm配置，连接mysql数据库
	global.Db = core.InitGorm()

	//迁移表结构
	//控制台输入 go run main.go -db
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	} else {
		//global.Log.Errorln("迁移表结构失败")
	}

	router := router.InitRouter()

	global.Log.Infoln("server start success is:", global.Config.System.Addr())

	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Errorln("server run err:", err.Error())
	}
}
