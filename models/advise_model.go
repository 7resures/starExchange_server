package models

import "gorm.io/gorm"

type Advise struct {
	gorm.Model
	AdviseId      int    `json:"adviseId" gorm:"primaryKey;autoIncrement"`
	AdviseContent string `json:"adviseContent"`
	UserId        int    `json:"userId"`
}
