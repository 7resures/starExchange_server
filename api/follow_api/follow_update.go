package follow_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type FollowUpdateRes struct {
	UserId      uint `json:"userId" form:"userId"`
	ProductId   uint `json:"productId" form:"productId"`
	IsConcerned bool `json:"isConcerned" form:"isConcerned"`
}

func (FollowApi) FollowUpdate(c *gin.Context) {
	req := FollowUpdateRes{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("参数绑定失败", c)
		return
	}

	//更新原商品状态
	query := global.Db.Model(&models.Product{})
	// 如果关注
	fmt.Println(req)
	if req.IsConcerned == true {
		result1 := query.Model(&models.Product{}).Where("product_id = ?", req.ProductId).Update("product_store", models.Product{}.ProductStore+1)
		result2 := global.Db.Model(&models.ProductsCollection{}).Create(&models.ProductsCollection{
			ProductId: req.ProductId,
			UserId:    req.UserId,
		})
		fmt.Println(result1.RowsAffected, result2.RowsAffected)
		if result1.RowsAffected == 0 || result2.RowsAffected == 0 {
			res.FailWithMessage("关注商品失败", c)
			return
		} else {
			res.OkWithMessage("关注成功", c)
			return
		}
	} else {
		result1 := query.Model(&models.Product{}).Where("product_id = ?", req.ProductId).Update("product_store", models.Product{}.ProductStore-1)
		result2 := global.Db.Unscoped().Where("product_id = ? and user_id = ?", req.ProductId, req.UserId).Delete(&models.ProductsCollection{})
		if result1.RowsAffected == 0 || result2.RowsAffected == 0 {
			res.FailWithMessage("取消关注商品失败", c)
			return
		} else {
			res.OkWithMessage("取消关注成功", c)
			return
		}
	}
	res.Ok(c)
}
