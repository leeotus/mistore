package models

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	mathRand "math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/google/uuid"
)

/**
 * @brief 获取当前时间的时间戳
 */
func TimeStamp() int64 {
	return time.Now().Unix()
}

/**
 * @brief 当前当前时间的时间戳,纳秒级别
 */
func TimeStampNano() int64 {
	return time.Now().UnixNano()
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

// @brief 字符串转换为浮点数
func Str2Float(str string) (float64, error) {
	n, err := strconv.ParseFloat(str, 64)
	return n, err
}

// @brief 字符串转换为HTML页面
func Str2Html(str string) template.HTML {
	return template.HTML(str)
}

/**
 * @brief int类型转换为string类型
 */
func Int2Str(num int) string {
	str := strconv.Itoa(num)
	return str
}

// @brief 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// Substr截取字符串
func Substr(str string, start int, end int) string {
	rs := []rune(str)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = 0
	}

	if end < 0 {
		end = rl
	}
	if end > rl {
		end = rl
	}
	if start > end {
		start, end = end, start
	}
	return string(rs[start:end])
}

// @brief 生成随机数
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := mathRand.Intn(10)
		str += Int2Str(current)
	}
	return str
}

func FormatAttr(str string) string {

	tempSlice := strings.Split(str, "\n")
	var tempStr string
	for _, v := range tempSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		tempStr += string(output)
	}
	return tempStr
}
