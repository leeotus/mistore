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

func RedisCoreInit(addr, port, pwd string, db, maxpool, minIdleConns int, maxConnAge, idleTimeout time.Duration) {
	host := fmt.Sprintf("%s:%s", addr, port)
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     host, // Redis地址
		Password: pwd,  // 密码
		DB:       db,   // 默认采用的数据库编号

		// 连接池配置
		PoolSize:     maxpool,      // 最大活跃连接数
		MinIdleConns: minIdleConns, // 最小空闲连接数
		MaxConnAge:   maxConnAge,   // 连接的最大存活时间
		IdleTimeout:  idleTimeout,  // 空闲连接的超时时间
	})

	// 测试连接
	ctx := context.Background()
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}
}
