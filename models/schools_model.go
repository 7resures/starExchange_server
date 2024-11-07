package models

import "gorm.io/gorm"

type School struct {
	gorm.Model
	SchoolID   uint   `json:"school_id" gorm:"primaryKey;"`
	SchoolName string `json:"school_name"`
	//Users      []User `gorm:"foreignKey:UserID"`
}
