package models

type Product struct {
	UserId             uint    `json:"userId" form:"userId"` // 外键，关联 User 表的 Id
	ProductId          uint    `json:"productId" form:"productId" gorm:"primary_key"`
	ProductName        string  `json:"productName" form:"productName"`
	ProductDescription string  `json:"productDescription" form:"productDescription"`
	ProductPrice       float64 `json:"productPrice" form:"productPrice"`
	ContactWeChat      string  `json:"contactWeChat" form:"contactWeChat"`
	ContactPhone       string  `json:"contactPhone" form:"contactPhone"`
	ContactQQ          string  `json:"contactQQ" form:"contactQQ"`
	ProductStatus      uint    `json:"productStatus" form:"productStatus"`
	ProductViews       uint    `json:"productViews" form:"productViews"`
	ProductStore       uint    `json:"productStore" form:"productStore"`
	Image              []Image `gorm:"foreignKey:ProductId;reference:ProductId"`
	Tags               []Tag   `gorm:"many2many:products_tags;"`
}
