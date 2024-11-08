package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductId          uint    `json:"product_id"`
	UserID             uint    `json:"user_id"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
	ContactWeChat      string  `json:"contact_we_chat"`
	ContactPhone       string  `json:"contact_phone"`
	ContactQQ          string  `json:"contact_qq"`
	ProductStatus      string  `json:"product_status"`
	Images             []Image `gorm:"foreignKey:ProductID"`
	Tags               []Tag   `gorm:"many2many:product_tags"`
}
