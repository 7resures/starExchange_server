package global

import (
	"EStarExchange/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Log    *logrus.Logger
	Rdb    *redis.Client
)

type UploadImgInfo struct {
	Status  string `json:"status"`
	Size    int64  `json:"size"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var (
	WhiteSuffix = []string{
		"jpg",
		"jpeg",
		"png",
		"svg",
		"ico",
		"webp",
	}
)
