package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

type UpdateGoodRequest struct {
	UserId             uint     `json:"userId" form:"userId"`
	ProductId          uint     `json:"productId" form:"productId"`
	ProductName        string   `json:"productName" form:"productName"`
	ProductDescription string   `json:"productDescription" form:"productDescription"`
	ProductPrice       float64  `json:"productPrice" form:"productPrice"`
	ContactWeChat      string   `json:"contactWeChat" form:"contactWeChat"`
	ContactPhone       string   `json:"contactPhone" form:"contactPhone"`
	ContactQQ          string   `json:"contactQQ" form:"contactQQ"`
	OriginImages       []string `json:"originImages" form:"originImages"`
}

func (GoodApi) GoodsInfoUpdate(c *gin.Context) {
	req := UpdateGoodRequest{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("数据绑定错误", c)
		return
	}
	fmt.Println("req is:", req)

	query := global.Db.Model(&models.Product{}).Where("product_id = ?", req.ProductId)

	//更新关注信息
	result := query.Updates(&models.Product{
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductPrice:       req.ProductPrice,
		ContactWeChat:      req.ContactWeChat,
		ContactPhone:       req.ContactPhone,
		ContactQQ:          req.ContactQQ,
	})
	if result.RowsAffected == 0 {
		res.FailWithMessage("更新商品信息失败", c)
		return
	}
	//删除云上图片
	for _, url := range req.OriginImages {
		index := strings.Index(url, "images/")
		if index != -1 {
			FileName := url[index+len("images/"):]
			basePath := filepath.Join("uploadFile", "image", FileName)

			// 检查文件是否存在
			if _, err := os.Stat(basePath); os.IsNotExist(err) {
				res.FailWithMessage(err.Error(), c)
				return
			}

			// 删除文件
			if err := os.Remove(basePath); err != nil {
				res.FailWithMessage(err.Error(), c)
				return
			}
		}
	}
	//删除数据库图片信息
	global.Db.Unscoped().Where("product_id = ?", req.ProductId).Delete(&models.Image{})

	//更新商品信息
	res.OkWithMessage("更新成功", c)
}
