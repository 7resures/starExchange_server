package image_api

import (
	"EStarExchange/global"
	"EStarExchange/router/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

type ImageRes struct {
	Status    bool   `json:"status"`
	ImageId   int    `json:"imageId"`
	ImageName string `json:"imageName"`
	Message   string `json:"message"`
}

func (ImageApi) ImagesUplpoad(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", c)
	}

	FileRes := []ImageRes{}

	//判断路径是否存在，不存在则创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	for index, file := range fileList {
		FilePath := path.Join(basePath, file.Filename)
		FileSize := float64(file.Size) / float64((1024 * 1024))
		if FileSize >= global.Config.Upload.Size {
			FileRes = append(FileRes, ImageRes{
				Status:    false,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   fmt.Sprintf("文件大小超出限制,文件大小限制为:%.1f MB,当前文件大小为 %.1f MB", global.Config.Upload.Size, FileSize),
			})
			continue
		}
		err := c.SaveUploadedFile(file, FilePath)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			FileRes = append(FileRes, ImageRes{
				Status:    false,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   err.Error(),
			})
		} else {
			FileRes = append(FileRes, ImageRes{
				Status:    true,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   "上传成功",
			})
		}

	}
	res.OkWithData(FileRes, c)
}
