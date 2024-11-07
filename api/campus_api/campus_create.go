package campus_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type campusVerify struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func (CampusApi) CreateCampus(c *gin.Context) {
	var req []models.School
	if err := c.ShouldBindJSON(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(req)
	var resp []campusVerify
	for _, item := range req {
		result := global.Db.Where("school_id = ? OR school_name = ?", item.SchoolID, strings.TrimSpace(item.SchoolName)).Find(&models.School{})
		if result.RowsAffected <= 0 {
			if err := global.Db.Create(&item).Error; err != nil {
				res.FailWithMessage(fmt.Sprintf("Failed to insert data: %v", err), c)
				return
			}
			resp = append(resp, campusVerify{
				Id:     item.SchoolID,
				Name:   item.SchoolName,
				Status: true,
			})
		} else {
			resp = append(resp, campusVerify{
				Id:     item.SchoolID,
				Name:   item.SchoolName,
				Status: false,
			})
		}
	}
	res.OkWithData(resp, c)
}
