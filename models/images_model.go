package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageID   uint   `json:"image_id" gorm:"primaryKey;"`
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}
