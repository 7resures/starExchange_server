package models

type User struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Identity    uint      `json:"identity" form:"identity"`
	Username    string    `json:"username" form:"username"`
	Nickname    string    `json:"nickname" form:"nickname"`
	Password    string    `json:"password" form:"password"`
	CampusName  string    `json:"campusName" form:"campusName"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	WeChat      string    `json:"wechat" form:"wechat"`
	WeChatId    string    `json:"wechatId" form:"wechatId"`
	Authority   uint      `json:"authority" form:"authority"`
	QQ          string    `json:"qq" form:"qq"`
	AvatarUrl   string    `json:"avatar" form:"avatar"`
	Product     []Product `gorm:"foreignKey:UserId;reference:Id"`
	Advise      []Advise  `gorm:"foreignKey:UserId;reference:Id"`
}
