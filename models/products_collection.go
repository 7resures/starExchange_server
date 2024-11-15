package models

import "gorm.io/gorm"

type ProductsCollection struct {
	gorm.Model
	ProductId uint `json:"productId" form:"productId"`
	UserId    uint `json:"userId" form:"userId"`
}
