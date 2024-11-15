package tag_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"github.com/gin-gonic/gin"
)

type TagList struct {
	Tags []string `json:"tags"`
}
type Res struct {
	Label string
	Info  string
}

func (TagApi) TagCreate(c *gin.Context) {
	var req TagList
	if err := c.BindJSON(&req); err != nil {
		res.FailWithMessage("获取参数失败", c)
		return
	}
	var resList = []Res{}
	for _, item := range req.Tags {
		result := global.Db.Where("tag_name = ?", item).Find(&models.Tag{})

		if result.RowsAffected != 0 {
			resList = append(resList, Res{
				Label: item,
				Info:  "该标签已存在",
			})
			continue
		}
		global.Db.Create(&models.Tag{
			TagName: item,
		})
		resList = append(resList, Res{
			Label: item,
			Info:  "标签创建成功",
		})

	}
	res.OkWithData(resList, c)
}
