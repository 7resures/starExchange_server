package login_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (LoginApi) RegisterHandler(c *gin.Context) {
	var req models.User

	// 使用 ShouldBindQuery 绑定查询参数
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithMessage("注册失败,请填写必要信息", c)
		return
	}
	// 检查用户是否存在，将用户信息插入到数据库中
	username := req.Username
	fmt.Println(req.Password)
	req.Password = utils.MD5([]byte(req.Password))
	fmt.Println(req.Password)
	checkUsername := global.Db.Where("username = ?", username).First(&models.User{})
	if checkUsername.RowsAffected > 0 {
		res.FailWithMessage(fmt.Sprintf("该用户 %s 已经存在", username), c)
		return
	} else {
		err := global.Db.Create(&req).Error
		if err != nil {
			res.FailWithMessage("网络错误,请重试!", c)
			return
		}
	}
	// 返回注册成功的响应
	res.OkWithMessage(fmt.Sprintf("用户 %s 注册成功", username), c)
}
