package user_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserPwdUpdate(c *gin.Context) {
	req := models.User{}
	if err := c.BindJSON(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	hashSecret := utils.MD5([]byte(req.Password))
	result := global.Db.Where("username = ?", req.Username).Updates(models.User{
		Password: hashSecret,
	})
	if result.RowsAffected == 0 {
		res.FailWithMessage("修改密码失败,请不要输入一样的密码", c)
		return
	}
	res.OkWithMessage("修改成功,密码哈希为:"+hashSecret, c)
}
