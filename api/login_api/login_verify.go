package login_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"github.com/gin-gonic/gin"
)

// 登录请求参数的结构体
//type LoginRequest struct {
//	Username string `json:"username"`
//	Password string `json:"password"`
//	Campus   string `json:"campus"`
//	Nickname string `json:"nickname"`
//	Avatar   string `json:"avatar"`
//	Wechat   string `json:"wechat"`
//}

// LoginVerify 用户登录
// @Tag 登录
// @Summary 登录
// @Description 登录
// @Param data body LoginRequest true "登录"
// @Router /api/login [post]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (LoginApi) LoginVerify(c *gin.Context) {

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(body)) // 打印请求体，确保它是正确的 JSON 格式
	//这个东西会改变c的类型

	var req models.User

	// 解析请求参数
	if err := c.BindJSON(&req); err != nil {
		res.FailWithMessage("用户名或者密码未输入,请重试!", c)
		return
	}
	var user models.User
	// 如果是微信授权登录，走这里，因为如果是账号登录，一定不会有wechat字段，如果有wechat字段，一定是授权登录
	// 如果是微信授权登录，那么有两种可能
	// 1，该授权的微信已经绑定设置过校园信息
	// 2. 该授权的微信为绑定设置过校园信息
	if req.WeChatId != "" {
		result := global.Db.Where("we_chat_id = ?", req.WeChatId).Find(&user)
		if result.RowsAffected <= 0 {
			err := global.Db.Create(&req).Error
			if err != nil {
				res.FailWithMessage("网络错误", c)
				return
			}
			UserInfo := utils.JwtInfo{
				Username: req.WeChat,
				Role:     0,
			}
			token, err := utils.CreateToken(UserInfo)
			if err != nil {
				res.FailWithMessage(err.Error(), c)
				return
			}
			//返回登录成功的响应
			res.OkWithData(token, c)
			return
		} else {
			// 代表该用户已经创建过用户直接进行登录
			UserInfo := utils.JwtInfo{
				Username: req.WeChatId,
				Role:     0,
				//Role 0 代表未绑定校园信息的用户
				//Role 1 代表已绑定校园信息的用户
			}
			token, err := utils.CreateToken(UserInfo)
			if err != nil {
				res.FailWithMessage(err.Error(), c)
				return
			}
			//返回登录成功的响应
			res.OkWithData(token, c)
			return
		}
	}

	result := global.Db.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		res.FailWithMessage("该用户未注册,请重试!", c)
		return
	}
	md5Pwd := utils.MD5([]byte(req.Password))
	result = global.Db.Where("username = ? and password = ?", req.Username, md5Pwd).First(&user)
	if result.Error != nil {
		res.FailWithMessage("密码错误", c)
		return
	}
	result = global.Db.Where("username = ? and campus_name = ?", req.Username, req.CampusName).First(&user)
	if result.Error != nil {
		res.FailWithMessage("归属校园错误", c)
		return
	}

	UserInfo := utils.JwtInfo{
		Username: user.Username,
		Role:     1,
	}
	token, err := utils.CreateToken(UserInfo)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//返回登录成功的响应
	res.OkWithData(token, c)
}
