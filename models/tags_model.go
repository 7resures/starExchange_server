package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName  string    `json:"tag_name"`
	Products []Product `json:"products" gorm:"many2many:product_tags;"`
}
