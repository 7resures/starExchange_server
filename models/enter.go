package models

type Page struct {
	PageNumber int    `form:"page_number"`
	PageSize   int    `form:"page_size"`
	Key        string `form:"key"`
	Sort       string `form:"sort"`
}
