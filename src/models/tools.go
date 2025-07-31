package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/google/uuid"
)

/**
 * @brief md5加密函数
 * @param str 要加密的字符串
 */
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

/**
 * @brief 使用Google的uuid生成SessionID
 */
func GenerateSessionUUID() string {
	return uuid.New().String()
}

/**
 * @brief 使用标准库来生成指定长度的SessionID
 */
func GenerateSessionID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		// TODO: 这里应该可以做修改
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
