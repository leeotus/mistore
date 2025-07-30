package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

// @brief 连接数据库
func init() {
	var ctx = context.Background()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
}
