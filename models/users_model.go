package models

type User struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Identity    uint      `json:"identity" form:"identity"`
	Username    string    `json:"username" form:"username"`
	Nickname    string    `json:"nickname" form:"nickname"` //首次绑定微信的名称
	Password    string    `json:"password" form:"password"`
	CampusId    uint      `json:"campusId" form:"campusId"`
	CampusName  string    `json:"campusName" form:"campusName"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	WeChat      string    `json:"wechat" form:"wechat"`
	WeChatId    string    `json:"wechatId" form:"wechatId"`
	Authority   uint      `json:"authority" form:"authority"`
	QQ          string    `json:"qq" form:"qq"`
	AvatarUrl   string    `json:"avatar" form:"avatar"`
	Products    []Product `gorm:"foreignKey:ProductId;"`
	Advises     []Advise  `gorm:"foreignKey:AdviseId;"`
}
