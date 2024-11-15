package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	utils "EStarExchange/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type GoodId struct {
	Pid uint `json:"Pid" form:"Pid"`
}

func (GoodApi) GoodsPicUpload(c *gin.Context) {
	var req GoodId
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	id := req.Pid
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	files := form.File["images"]
	////判断路径是否存在，不存在则创建
	basePath := global.Config.Upload.Image
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	FileRes := []models.ImageRes{}

	for index, file := range files {
		fileOjb, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := ioutil.ReadAll(fileOjb)
		md5String := utils.MD5(byteData)
		timestamp := time.Now().Unix()
		newFileName := fmt.Sprintf("%s_%d_%d%s", md5String, index, timestamp, path.Ext(file.Filename))
		FilePath := path.Join(basePath, newFileName)

		FileSize := float64(file.Size) / float64((1024 * 1024))
		if FileSize >= global.Config.Upload.Size {
			FileRes = append(FileRes, models.ImageRes{
				Status:    false,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   fmt.Sprintf("文件大小超出限制,文件大小限制为:%.1f MB,当前文件大小为 %.1f MB", global.Config.Upload.Size, FileSize),
			})
			continue
		}
		err = c.SaveUploadedFile(file, FilePath)
		if err != nil {
			FileRes = append(FileRes, models.ImageRes{
				Status:    false,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   err.Error(),
			})
			res.FailWithMessage(err.Error(), c)
			return
		} else {
			imgurl := fmt.Sprintf("http://%s:%d/images/%s", global.Config.System.Host, global.Config.System.Port, newFileName)
			result := global.Db.Create(&models.Image{
				ProductId: id,
				ImageURL:  imgurl,
			})
			if result.RowsAffected == 0 {
				FileRes = append(FileRes, models.ImageRes{
					Status:    true,
					ImageId:   index,
					ImageName: file.Filename,
					Message:   "上传图片失败",
				})
				res.FailWithData(FileRes, c)
				return
			}
			FileRes = append(FileRes, models.ImageRes{
				Status:    true,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   "上传成功",
			})
		}
	}

	res.OkWithData(FileRes, c)
}
