package db

/**
 * @brief Redis数据库连接初始化
 * @author leeotus
 * @email leeotus@163.com
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/**
 * @brief Redis线程池
 */

var RedisDB *redis.Client
var ctx = context.Background()

type RDBCache struct{}

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
	// ctx := context.Background()
	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis连接失败: %v", err))
	}
}

/**
 * @brief 向Redis里设置数据
 * @param key 键
 * @param value 键对应的值
 * @param expiration 该键值对的过期时间
 */
func (RDBCache) Set(key string, value any, expiration int) {
	bytes, _ := json.Marshal(value)
	RedisDB.Set(ctx, key, string(bytes), time.Second*time.Duration(expiration))
}

/**
 * @brief Set函数,带context版
 */
func (RDBCache) SetWithContext(ic context.Context, key string, value any, expiration int) {
	bytes, _ := json.Marshal(value)
	RedisDB.Set(ic, key, string(bytes), time.Second*time.Duration(expiration))
}

/**
 * @brief 企图获取保存在Redis里的数据,以输入"key"作为键值
 * @param obj 若获取到数据,则保存在obj中
 * @return bool 成功获取数据返回true,否则返回false
 */
func (RDBCache) Get(key string, obj any) bool {
	redisStr, err1 := RedisDB.Get(ctx, key).Result()
	if err1 == nil && redisStr != "" {
		err2 := json.Unmarshal([]byte(redisStr), obj)
		return err2 == nil
	}
	return false
}

/**
 * @brief Get函数,带context版
 */
func (RDBCache) GetWithCache(ic context.Context, key string, obj any) bool {
	redisStr, err1 := RedisDB.Get(ic, key).Result()
	if err1 == nil && redisStr != "" {
		err2 := json.Unmarshal([]byte(redisStr), obj)
		return err2 == nil
	}
	return false
}

var RedisCache = &RDBCache{}
