package models

type Tag struct {
	Id       uint      `gorm:"primary_key;auto_increment" json:"id"`
	TagName  string    `json:"tag_name"`
	Products []Product `gorm:"many2many:products_tags;"`
}
