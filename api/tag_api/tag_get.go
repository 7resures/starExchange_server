package tag_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagGet(c *gin.Context) {
	var tags []models.Tag
	result := global.Db.Model(&models.Tag{}).Find(&tags)
	if result.RowsAffected == 0 {
		res.FailWithMessage("没有找到相关标签", c)
		return
	}
	res.OkWithData(tags, c)
}
