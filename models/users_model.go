package models

type User struct {
	ID                uint               `json:"id" gorm:"primaryKey;"`
	UserID            uint               `json:"user_id"`
	Identity          uint               `json:"identity"`
	Username          string             `json:"username" form:"username"`
	Nickname          string             `json:"nickname" form:"nickname"`
	Password          string             `json:"password" form:"password"`
	CampusID          uint               `json:"campus_id"`
	CampusName        string             `json:"campus_name"`
	PhoneNumber       string             `json:"phone_number"`
	WeChat            string             `json:"wechat"`
	Authority         uint               `json:"authority"`
	QQ                string             `json:"qq"`
	AvatarURL         string             `json:"avatar"`
	Products          []Product          `gorm:"foreignKey:ProductID;"`
	SecurityQuestions []SecurityQuestion `gorm:"foreignKey:SecurityID;"`
}
