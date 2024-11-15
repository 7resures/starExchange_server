package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageID   uint   `json:"imageId"`
	ProductId uint   `json:"productId"`
	ImageURL  string `json:"imageUrl"`
}
