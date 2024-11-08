package login_api

import (
	"EStarExchange/global"
	"EStarExchange/router/res"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type tokenRes struct {
	Code string `form:"code"`
}

type WxLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid,omitempty"`
	ErrCode    int    `json:"errcode,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
}

func (LoginApi) TokenGet(c *gin.Context) {
	req := tokenRes{}
	if err := c.ShouldBindQuery(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fmt.Println("req is :", req)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + string(global.Config.Wx.Appid) + "&secret=" + string(global.Config.Wx.Secret) + "&js_code=" + req.Code + "&grant_type=authorization_code"
	fmt.Println(url)
	// 发起 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var wxResp WxLoginResponse
	if err := json.Unmarshal(body, &wxResp); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithData(wxResp, c)
}
