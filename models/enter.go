package models

type Page struct {
	PageNumber int    `form:"page_number"`
	PageSize   int    `form:"page_size"`
	Key        string `form:"key"`
	Sort       string `form:"sort"`
}

type ImageRes struct {
	Status    bool   `json:"status"`
	ImageId   int    `json:"imageId"`
	ImageName string `json:"imageName"`
	Message   string `json:"message"`
}
