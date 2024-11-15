package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type DeleteInfo struct {
	UserId    uint `json:"userId" form:"userId"`
	ProductId uint `json:"productId" form:"productId"`
}

func (GoodApi) GoodsInfoDelete(c *gin.Context) {
	//body, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// 如果读取失败，返回错误
	//	c.JSON(400, gin.H{
	//		"error": "Failed to read body",
	//	})
	//	return
	//}
	//
	//// 打印原始请求数据
	//fmt.Println("Raw Request Body:", string(body))
	req := DeleteInfo{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("删除商品信息有误!", c)
		return
	}
	fmt.Println(req)
	//删除收藏信息
	result1 := global.Db.Where("product_id = ?  and user_id = ?", req.ProductId, req.UserId).Delete(&models.ProductsCollection{})
	if result1.Error != nil {
		res.FailWithMessage("收藏信息有问题!", c)
		return
	}
	//删除商品图片信息
	result2 := global.Db.Unscoped().Where("product_id = ?", req.ProductId).Delete(&models.Image{})
	if result2.Error != nil {
		res.FailWithMessage("图片信息有问题!", c)
		return
	}
	//删除商品标签信息
	result3 := global.Db.Where("product_id = ?", req.ProductId).Delete(&models.ProductsTags{})
	if result3.Error != nil {
		res.FailWithMessage("商品标签有问题!", c)
		return
	}
	//删除商品信息
	result4 := global.Db.Where("product_id = ?", req.ProductId).Delete(&models.Product{})
	if result4.Error != nil {
		res.FailWithMessage("商品删除失败", c)
		return
	}

	//删除商品收藏信息

	res.Ok(c)
}
