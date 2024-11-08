package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	SchoolID       uint   `json:"schoolId" form:"schoolId" gorm:"primaryKey;"`
	SchoolName     string `json:"schoolName" form:"schoolName"`
	SchoolProvince string `json:"schoolProvince" form:"schoolProvince"`
	//Users      []User `gorm:"foreignKey:UserID"`
}
