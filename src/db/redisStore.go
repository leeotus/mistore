package db

import (
	"context"
	"fmt"
	"time"
)

/**
 * @brief 在这里去实现base64Captcha.Store接口,使得之后可以直接用来适配
 */

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct{}

// @brief 设置captcha
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := RedisDB.Set(ctx, key, value, time.Minute*2).Err()
	return err
}

/**
 * @brief 获取captcha验证码
 * @param id 需要获取其值的key
 * @param clear 是否需要删除数据库里的数据
 */
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDB.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := RedisDB.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

/**
 * @brief 验证captcha
 */
func (r RedisStore) Verify(id string, ans string, clear bool) bool {
	v := r.Get(id, clear)
	fmt.Println("key:" + id + ";value:" + v + ";ans:" + ans)
	return v == ans
}
