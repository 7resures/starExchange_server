package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type StatusInfo struct {
	ProductId     uint `json:"productId" form:"productId"`
	ProductStatus uint `json:"productStatus" form:"productStatus"`
}

func (GoodApi) GoodStatusUpdate(c *gin.Context) {
	req := StatusInfo{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(req)
	result := global.Db.Model(&models.Product{}).Where("product_id = ? ", req.ProductId).Update("product_status", req.ProductStatus)
	if result.RowsAffected == 0 {
		res.FailWithMessage("未找到该商品", c)
		return
	}
	res.Ok(c)
}
