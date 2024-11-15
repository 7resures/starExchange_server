package models

type ProductsTags struct {
	Id        uint `json:"id" gorm:"primary_key;auto_increment"`
	TagId     uint `json:"tag" form:"tag" `
	ProductId uint `json:"productId" form:"productId"`
}
