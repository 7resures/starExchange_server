package flag

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"fmt"
)

func MakeMigrations() {
	fmt.Println("迁移执行中...")
	global.Db.Migrator().AutoMigrate(&models.User{}, &models.Product{}, &models.Image{}, &models.Tag{}, &models.School{}, &models.Advise{}, &models.ProductsTags{}, &models.ProductsCollection{})
	fmt.Println("迁移执行成功")
}
