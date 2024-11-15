package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	SchoolName     string `json:"schoolName" form:"schoolName"`
	SchoolProvince string `json:"schoolProvince" form:"schoolProvince"`
}
