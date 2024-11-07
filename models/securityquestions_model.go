package models

import "gorm.io/gorm"

type SecurityQuestion struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primaryKey;"`
	SecurityID uint   `json:"security_id"`
	UserID     uint   `json:"user_id"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
}
