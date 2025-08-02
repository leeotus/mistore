package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

/**
 * @brief 获取当前时间的时间戳
 */
func TimeStamp() int64 {
	return time.Now().Unix()
}

// 时间戳转换成日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

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

/**
 * @brief string类型的数据转换为int类型
 */
func Str2Int(str string) (int, error) {
	num, err := strconv.Atoi(str)
	return num, err
}

/**
 * @brief int类型转换为string类型
 */
func Int2Str(num int) string {
	str := strconv.Itoa(num)
	return str
}
