package user_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserInfoUpdate(c *gin.Context) {
	req := models.User{}
	if err := c.BindJSON(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if req.Username == "" {
		res.FailWithMessage("登录账号不能设置为空,请重试", c)
		return
	}

	confirm := global.Db.Where("username = ? or we_chat_id = ? and we_chat_id != '' ", req.Username, req.WeChatId).First(&models.User{})
	if confirm.RowsAffected <= 0 {
		res.FailWithMessage("非法用户", c)
		return
	}
	result := global.Db.Where("username = ? or we_chat_id = ? and we_chat_id != '' ", req.Username, req.WeChatId).Updates(models.User{
		Username:    req.Username,
		Nickname:    req.Nickname,
		CampusName:  req.CampusName,
		PhoneNumber: req.PhoneNumber,
		WeChat:      req.WeChat,
		QQ:          req.QQ,
		Authority:   req.Authority,
	})
	if result.RowsAffected <= 0 {
		global.Log.Errorln(result.Error)
		res.FailWithMessage("更新用户数据失败,登录账号不允许修改", c)
		return
	}
	res.OkWithMessage("更新成功", c)
}
