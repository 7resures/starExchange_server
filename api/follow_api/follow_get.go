package follow_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type FollowGetReq struct {
	UserId    uint `json:"userId" form:"userId"`
	ProductId uint `json:"productId" form:"productId"`
}

func (FollowApi) FollowGet(c *gin.Context) {
	req := FollowGetReq{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("非法用户Id", c)
		return
	}
	fmt.Println(req)
	result := global.Db.Where("user_id = ? and product_id = ?", req.UserId, req.ProductId).Find(&models.ProductsCollection{})
	if result.RowsAffected > 0 {
		res.OkWithData(true, c)
		return
	}
	res.OkWithData(false, c)

}
