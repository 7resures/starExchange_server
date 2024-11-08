package user_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"github.com/gin-gonic/gin"
)

// 用于查询用户信息
type UserTitleInfo struct {
	Token string `form:"token"`
}

func (UserApi) UserInfoGet(c *gin.Context) {
	userinfo := models.User{}
	req := UserTitleInfo{}
	//queryParams := c.Request.URL.Query()
	//fmt.Println("所有查询参数:", queryParams)
	if err := c.BindQuery(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	//解析token是否过期
	tokenInfo, err := utils.AnalysToken(req.Token)
	//过期则返回对应的信息
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	searchName := tokenInfo.Username
	//未过期则使用token解析后的信息返回当前用户所有的信息
	result := global.Db.Where("username = ? OR we_chat_id = ?", searchName, searchName).First(&userinfo)
	if result.RowsAffected == 0 {
		res.FailWithMessage("用户信息加载失败", c)
		return
	}
	// 这里将 `userinfo` 数据结构体直接作为第一个参数传递给 OkWithData 方法
	res.OkWithData(userinfo, c)
}
