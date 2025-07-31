package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @brief Redis线程池
 */

var RedisDB *redis.Client

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis地址
		Password: "",               // 密码
		DB:       0,                // 默认采用的数据库编号

		// 连接池配置
		PoolSize:     12,                // 最大活跃连接数
		MinIdleConns: 2,                 // 最小空闲连接数
		MaxConnAge:   300 * time.Second, // 连接的最大存活时间
		IdleTimeout:  60 * time.Second,  // 空闲连接的超时时间
	})

	// 测试连接
	ctx := context.Background()
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}
}
