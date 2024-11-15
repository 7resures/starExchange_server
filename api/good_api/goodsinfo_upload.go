package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"github.com/gin-gonic/gin"
)

type Good struct {
	ProductId          uint     `json:"productId" form:"productId"`
	UserID             uint     `json:"userId" form:"userId"`
	ProductName        string   `json:"productName" form:"productName"`
	ProductDescription string   `json:"productDescription" form:"productDescription"`
	ProductPrice       float64  `json:"productPrice" form:"productPrice"`
	ContactWeChat      string   `json:"contactWeChat" form:"contactWeChat"`
	ContactPhone       string   `json:"contactPhone" form:"contactPhone"`
	ContactQQ          string   `json:"contactQQ" form:"contactQQ"`
	ProductStatus      uint     `json:"productStatus" form:"productStatus"`
	ProductViews       uint     `json:"productViews" form:"productViews"`
	ProductStore       uint     `json:"productStore" form:"productStore"`
	Tags               []string `json:"tags" form:"tags"`
}

func (GoodApi) GoodsInfoUpload(c *gin.Context) {
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

	req := Good{}
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("无效数据", c)
		return
	}

	sf, err := utils.NewSnowflake(1)
	if err != nil {
		global.Log.Error(err.Error())
	}
	Pid := sf.Generate()
	result := global.Db.Create(&models.Product{
		ProductId:          uint(Pid),
		UserId:             req.UserID,
		ProductName:        req.ProductName,
		ProductDescription: req.ProductDescription,
		ProductPrice:       req.ProductPrice,
		ContactWeChat:      req.ContactWeChat,
		ContactPhone:       req.ContactPhone,
		ContactQQ:          req.ContactQQ,
		ProductStatus:      req.ProductStatus,
		ProductViews:       req.ProductViews,
		ProductStore:       req.ProductStore,
	})
	if result.RowsAffected == 0 {
		res.FailWithMessage("商品上传失败", c)
		return
	}

	for _, item := range req.Tags {
		var Tag models.Tag
		result := global.Db.Where("tag_name = ?", item).Find(&Tag)
		if result.RowsAffected == 0 {
			res.FailWithMessage("商品标签不存在", c)
			return
		}
		result = global.Db.Create(&models.ProductsTags{
			TagId:     Tag.Id,
			ProductId: uint(Pid),
		})
		if result.RowsAffected == 0 {
			res.FailWithMessage("标签插入失败", c)
			return
		}
	}

	res.OkWithData(uint(Pid), c)
}
