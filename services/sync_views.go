package services

import (
	"EStarExchange/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// 定时任务将商品浏览次数写入数据库
func SyncBrowseCountsToDatabase() {
	ticker := time.NewTicker(10 * time.Second) // 每 10 秒执行一次
	defer ticker.Stop()

	for range ticker.C {
		// 获取所有商品的浏览次数
		productIDs := []string{"product1", "product2", "product3"} // 假设这是我们要同步的商品ID

		for _, productID := range productIDs {
			// 获取商品浏览次数
			browseCount, err := global.Rdb.Get(context.Background(), productID).Result()
			if err == redis.Nil {
				// 商品的浏览次数没有被设置过
				browseCount = "0"
			} else if err != nil {
				log.Printf("failed to get browse count for %s: %v", productID, err)
				continue
			}

			// 将浏览次数写入数据库
			fmt.Println(browseCount)
		}
	}
}
