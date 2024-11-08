package advise_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

func (AdviseApi) AdviseCreate(c *gin.Context) {
	req := models.Advise{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	result := global.Db.Create(&req)
	if result.RowsAffected == 0 {
		res.FailWithMessage("提交失败", c)
		return
	}
	res.OkWithMessage("提交成功", c)
}
