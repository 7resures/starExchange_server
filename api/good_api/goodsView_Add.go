package good_api

import (
	"EStarExchange/global"
	"EStarExchange/models"
	"EStarExchange/router/res"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
)

var (
	mutex      sync.RWMutex // 读写锁
	browseLock sync.Map     // 用来存储每个商品的锁，避免重复加锁
)

type GoodsView struct {
	ProductId uint `json:"productId" form:"productId"`
}

func (GoodApi) GoodsAddView(c *gin.Context) {
	req := GoodsView{}
	resCount := 0
	if err := c.ShouldBind(&req); err != nil {
		res.FailWithMessage("商品id无效", c)
		return
	}

	// 检查商品ID是否存在
	exists, err := global.Rdb.SIsMember(context.Background(), "allProductIds", req.ProductId).Result()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	// 商品不存在
	if !exists {
		var productViews int
		result := global.Db.Model(&models.Product{}).
			Where("product_id = ?", req.ProductId).
			Pluck("product_views", &productViews)
		if result.Error != nil {
			res.FailWithMessage("商品id错误", c)
			return
		}

		// 将商品ID加入到 allProductIds 集合
		err := global.Rdb.SAdd(context.Background(), "allProductIds", req.ProductId).Err()
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		// 将该 productId 和值存入 Redis Hash
		err = global.Rdb.HSet(context.Background(), "productValues", req.ProductId, productViews+1).Err()
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		resCount = productViews + 1
	} else {
		// 获取 productId 对应的 product_views
		value, err := global.Rdb.HGet(context.Background(), "productValues", strconv.Itoa(int(req.ProductId))).Result()
		if err == redis.Nil {
			// 如果找不到该商品的浏览次数
			res.FailWithMessage(err.Error(), c)
			return
		} else if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		// 获取当前浏览次数并增加
		currentValue, _ := strconv.Atoi(value)
		newValue := currentValue + 1

		// 使用 Redis 锁确保并发安全
		productIdKey := fmt.Sprintf("lock:%d", req.ProductId)
		locked, err := global.Rdb.SetNX(context.Background(), productIdKey, "locked", 10*time.Second).Result()
		if err != nil {
			res.FailWithMessage("无法获取锁", c)
			return
		}

		if !locked {
			// 如果无法获取锁，说明有其他请求在修改该商品的浏览次数
			res.FailWithMessage("商品正在被访问，请稍后重试", c)
			return
		}

		// 设置新的浏览次数
		err = global.Rdb.HSet(context.Background(), "productValues", strconv.Itoa(int(req.ProductId)), newValue).Err()
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}

		// 解锁
		global.Rdb.Del(context.Background(), productIdKey)

		resCount = newValue + 1
	}

	// 返回响应
	res.OkWithData(resCount, c)
}

// 定时任务：每 60 分钟写回数据库
func init() {
	go func() {
		ticker := time.NewTicker(30 * time.Minute) // 每60分钟执行一次
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				writeBackToDatabase()
			}
		}
	}()
}

// 将 Redis 中的商品浏览次数写回数据库
func writeBackToDatabase() {
	// 获取 Redis 中所有商品的浏览次数
	productValues, err := global.Rdb.HGetAll(context.Background(), "productValues").Result()
	if err != nil {
		fmt.Println("无法获取 Redis 中的商品浏览次数：", err)
		return
	}

	// 将 Redis 中的商品浏览次数写回数据库
	for productIdStr, viewsStr := range productValues {
		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			fmt.Println("商品ID转换错误:", err)
			continue
		}
		views, err := strconv.Atoi(viewsStr)
		if err != nil {
			fmt.Println("商品浏览次数转换错误:", err)
			continue
		}

		// 更新数据库中的商品浏览次数
		err = global.Db.Model(&models.Product{}).
			Where("product_id = ?", productId).
			Update("product_views", views).Error
		if err != nil {
			fmt.Println("写回数据库失败:", err)
		}
	}

}
