package campus_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

func (CampusApi) CampusGet(c *gin.Context) {
	var CampusInfo []models.School

	// 查询所有学校记录
	result := global.Db.Find(&CampusInfo)

	// 如果查询出错，返回服务器错误
	if result.Error != nil {
		res.FailWithMessage("服务器错误", c)
		return
	}

	// 查询成功，收集学校名称
	var CampusName []string
	for _, item := range CampusInfo {
		CampusName = append(CampusName, item.SchoolName)
	}

	// 返回学校名称列表
	res.OkWithData(CampusName, c)
}
