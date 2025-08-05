package utils

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// NOTE: 这里需要改成自己的名称
// TODO: 之后可以写入到.ini配置文件里
const DEFAULT_BUCKET_NAME = "leeotus-mistore"

var OSSClient *oss.Client
var OSSBucket *oss.Bucket

/**
 * @brief 使用ini文件来初始化
 * @note 需要在对应的.ini文件里保存AccessKey等信息,但是考虑到安全性的问题,实际使用不推荐这种使用方式
 * @todo 一种方法是:在/admin界面里可以选择是开启oss存储,并在那里存放AccessKey等信息,admin管理员自己输入对应的信息
 */
func OOSClientInitFromIni(endpoint string, accessKey string, accessSecret string) {
	fmt.Println("endpoint:", endpoint)
	fmt.Println("accessKey:", accessKey)
	fmt.Println("accessSecret:", accessSecret)
	var err1 error
	OSSClient, err1 = oss.New(endpoint, accessKey, accessSecret)
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(-1)
	}
	// lsRes, err2 := OOSClient.Bucket(DEFAULT_BUCKET_NAME)
	isExist, err2 := OSSClient.IsBucketExist(DEFAULT_BUCKET_NAME)
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(-1)
	}
	if !isExist {
		err3 := OSSClient.CreateBucket(DEFAULT_BUCKET_NAME)
		if err3 != nil {
			fmt.Println("创建bucket失败,考虑手动创建并修改'DEFAULT_BUCKET_NAME'的值:" + err3.Error())
			os.Exit(-1)
		}
	}
	var err4 error
	OSSBucket, err4 = OSSClient.Bucket(DEFAULT_BUCKET_NAME)
	if err4 != nil {
		fmt.Println(err4)
		os.Exit(-1)
	}
}
