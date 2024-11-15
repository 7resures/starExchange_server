package core

import (
	"EStarExchange/global"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	rdbConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdbConfig.GetAddr(),
		Password: rdbConfig.Password,
		DB:       db,
		PoolSize: rdbConfig.Pool_size,
	})

	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel() // 确保函数结束时取消上下文

	// Ping Redis 检查连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("Redis 连接失败: %v", err)
		return nil
	}

	logrus.Infof("Redis 连接成功: %s", rdbConfig.GetAddr())
	return rdb
}
