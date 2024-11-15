package user_api

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

type Userid struct {
	id     int    `form:"id"`
	images string `form:"images"`
}

func (UserApi) AvatarUpdate(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	// 获取文本字段 id
	id := form.Value["id"]
	if len(id) == 0 {
		res.FailWithMessage("缺少id", c)
		return
	}

	fileList, ok := form.File["imageKey"]
	if !ok {
		res.FailWithMessage("不存在文件", c)
		return
	}

	FileRes := []models.ImageRes{}

	//判断路径是否存在，不存在则创建
	basePath := global.Config.Upload.Avatar
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	for index, file := range fileList {

		//修改当前文件的文件名
		fileOjb, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := ioutil.ReadAll(fileOjb)
		md5String := utils.MD5(byteData)
		timestamp := time.Now().Unix()
		newFileName := fmt.Sprintf("%s_%s_%d%s", md5String, id[0], timestamp, path.Ext(file.Filename))
		FilePath := path.Join(basePath, newFileName)

		// 判断是否超过上传文件的大小限制
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

		//保存文件
		err = c.SaveUploadedFile(file, FilePath)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			fmt.Println(file.Filename)
			FileRes = append(FileRes, models.ImageRes{
				Status:    false,
				ImageId:   index,
				ImageName: file.Filename,
				Message:   err.Error(),
			})
		} else {
			imgurl := fmt.Sprintf("http://%s:%d/avatars/%s", global.Config.System.Host, global.Config.System.Port, newFileName)
			fmt.Println(imgurl)
			result := global.Db.Where("id = ?", id).Updates(models.User{
				AvatarUrl: imgurl,
			})
			if result.RowsAffected == 0 {
				res.FailWithMessage(err.Error(), c)
				return
			}
			FileRes = append(FileRes, models.ImageRes{
				Status:    true,
				ImageId:   index,
				ImageName: fmt.Sprintln("images/" + FilePath),
				Message:   "上传成功",
			})
		}

	}
	res.OkWithData(FileRes, c)
}
