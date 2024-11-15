package models

import "gorm.io/gorm"

type Advise struct {
	gorm.Model
	AdviseContent string `json:"adviseContent"`
	UserId        uint   `json:"userId"`
}
