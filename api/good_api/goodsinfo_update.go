package good_api

import (
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

type UpdateGoodRequest struct {
	UserId             uint    `json:"userId" form:"userId"`
	ProductId          uint    `json:"productId" form:"productId"`
	IsConcern          bool    `json:"isConcern" form:"isConcern"`
	ProductName        string  `json:"productName" form:"productName"`
	ProductDescription string  `json:"productDescription" form:"productDescription"`
	ProductPrice       float64 `json:"productPrice" form:"productPrice"`
	ContactWeChat      string  `json:"contactWeChat" form:"contactWeChat"`
	ContactPhone       string  `json:"contactPhone" form:"contactPhone"`
	ContactQQ          string  `json:"contactQQ" form:"contactQQ"`
}

func (GoodApi) GoodsInfoUpdate(c *gin.Context) {
	req := UpdateGoodRequest{}
	if err := c.ShouldBind(req); err != nil {
		res.FailWithMessage("数据绑定错误", c)
	}

	//更新（删除）关注信息

	//更新商品信息

}
