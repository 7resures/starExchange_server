package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Request struct {
	UserId        uint   `json:"userId" form:"userId"`
	ProductId     uint   `json:"productId" form:"productId"`
	SearchContent string `json:"searchContent" form:"searchContent"`
	Concern       bool   `json:"concern" form:"concern"`
	Page          int    `json:"page" form:"page"`
	CampusName    string `json:"campusName" form:"campusName"`
	Type          uint   `json:"type" form:"type"`
	PageSize      int    `json:"pageSize" form:"pageSize"`
	Random        bool   `json:"random" form:"random"`
}

type GoodsInfo struct {
	UserId             uint    `json:"userId" form:"userId"` // 外键，关联 User 表的 Id
	ProductId          uint    `json:"productId" form:"productId" gorm:"primaryKey"`
	ProductName        string  `json:"productName" form:"productName"`
	ProductDescription string  `json:"productDescription" form:"productDescription"`
	ProductPrice       float64 `json:"productPrice" form:"productPrice"`
	ContactWeChat      string  `json:"contactWeChat" form:"contactWeChat"`
	ContactPhone       string  `json:"contactPhone" form:"contactPhone"`
	ContactQQ          string  `json:"contactQQ" form:"contactQQ"`
	ProductStatus      uint    `json:"productStatus" form:"productStatus"`
	ProductViews       uint    `json:"productViews" form:"productViews"`
	ProductStore       uint    `json:"productStore" form:"productStore"`
}

type ImageTags struct {
	Images []string `json:"images" form:"images"`
	Tags   []uint   `json:"tags" form:"tags"`
}

type GoodsWithImagesAndTags struct {
	GoodsInfo
	ImageTags
}

func (GoodApi) GoodsGet(c *gin.Context) {
	var req Request
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("参数绑定错误", c)
		return
	}

	// 定义返回的商品列表
	var goodRes []GoodsWithImagesAndTags
	var querys = global.Db.Model(&models.Product{})
	var query *gorm.DB
	if req.CampusName != "" {
		query = querys.Joins("JOIN users us ON us.id = products.user_id").
			Where("us.campus_name = ?", req.CampusName) // 基础查询对象
	} else {
		query = querys
	}

	// 1. 获取单条商品信息
	if req.ProductId != 0 {
		var singleGood GoodsWithImagesAndTags
		result := query.Where("product_id = ?", req.ProductId).First(&singleGood.GoodsInfo)
		if result.RowsAffected == 0 {
			res.FailWithMessage("未找到该商品", c)
			return
		}
		// 查询商品的图片和标签
		loadImagesAndTags(&singleGood.ImageTags, singleGood.GoodsInfo.ProductId)
		res.OkWithData(singleGood, c)
		return
	}

	//关注数据
	if req.Concern == true {
		query1 := query.Joins("JOIN products_collections pc ON pc.product_id = products.product_id").Where("pc.user_id = ?", req.UserId)
		result := query1.Find(&goodRes)
		if result.RowsAffected == 0 {
			res.FailWithMessage("暂无更多关注数据", c)
			return
		}
		for i := range goodRes {
			loadImagesAndTags(&goodRes[i].ImageTags, goodRes[i].GoodsInfo.ProductId)
		}
		res.OkWithData(goodRes, c)
		return
	}

	// 2. 获取某用户的所有商品信息
	if req.UserId != 0 {
		result := query.Where("user_id = ?", req.UserId).Find(&goodRes)
		if result.RowsAffected == 0 {
			res.FailWithMessage("该用户无商品数据", c)
			return
		}
		for i := range goodRes {
			loadImagesAndTags(&goodRes[i].ImageTags, goodRes[i].GoodsInfo.ProductId)
		}
		res.OkWithData(goodRes, c)
		return
	}

	if req.SearchContent != "" {
		result := query.Where("product_description LIKE ? or product_name LIKE ?", "%"+req.SearchContent+"%", "%"+req.SearchContent+"%").Find(&goodRes)
		if result.RowsAffected == 0 {
			res.FailWithMessage("该用户无商品数据", c)
			return
		}
		for i := range goodRes {
			loadImagesAndTags(&goodRes[i].ImageTags, goodRes[i].GoodsInfo.ProductId)
		}
		res.OkWithData(goodRes, c)
		return
	}

	// 3. 根据 type 联合查询 products_tags 表
	if req.Type != 0 {
		// 联合查询 products 和 products_tags 表
		query = query.Joins("JOIN products_tags pt ON pt.product_id = products.product_id").
			Where("pt.tag_id = ? ", req.Type)
	}

	// 4. 分页查询商品信息
	if req.Page > 0 && req.PageSize > 0 {
		offset := (req.Page - 1) * req.PageSize
		// 如果需要随机获取一页商品
		if req.Random {
			query = query.Order("RAND()")
		}
		result := query.Limit(req.PageSize).Offset(offset).Find(&goodRes)
		if result.RowsAffected == 0 {
			res.FailWithMessage("暂无更多数据", c)
			return
		}

		for i := range goodRes {
			loadImagesAndTags(&goodRes[i].ImageTags, goodRes[i].GoodsInfo.ProductId)
		}
		res.OkWithData(goodRes, c)
		return
	}

	// 5. 默认返回所有商品（不建议直接使用，除非明确需求）
	result := query.Find(&goodRes)
	if result.RowsAffected == 0 {
		res.FailWithMessage("暂无商品数据", c)
		return
	}

	for i := range goodRes {
		loadImagesAndTags(&goodRes[i].ImageTags, goodRes[i].GoodsInfo.ProductId)
	}
	res.OkWithData(goodRes, c)
}

// 辅助函数：加载商品的图片和标签
func loadImagesAndTags(imageTags *ImageTags, productId uint) {
	// 查询图片
	images := []models.Image{}
	global.Db.Where("product_id = ?", productId).Find(&images)
	for _, item1 := range images {
		imageTags.Images = append(imageTags.Images, item1.ImageURL)
	}

	// 查询标签
	tags := []models.ProductsTags{}
	global.Db.Where("product_id = ?", productId).Find(&tags)
	for _, item2 := range tags {
		imageTags.Tags = append(imageTags.Tags, item2.TagId)
	}
}
