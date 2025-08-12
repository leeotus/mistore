package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type UserCookie struct{}

const COOKIE_LIFETIME = 3600 * 24 * 30

// @brief 数据写入cookie
func (UserCookie) Set(c *gin.Context, key string, value any) {
	bytes, _ := json.Marshal(value)
	c.SetCookie(key, string(bytes), COOKIE_LIFETIME, "/", c.Request.Host, false, true)
}

// @brief 获取数据
func (UserCookie) Get(c *gin.Context, key string, obj any) bool {
	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		err2 := json.Unmarshal([]byte(valueStr), obj)
		return err2 == nil
	}
	return false
}

func (UserCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

var Cookie = &UserCookie{}
