package admin

/**
 * @brief 负责处理图片相关的函数,例如上传图片,裁剪图片
 */

import (
	"errors"
	"fmt"
	"mime/multipart"
	"mistore/src/models"
	"mistore/src/utils"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

const UPLOAD_DIR = "static/upload/"

var allowExtMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
	".gif":  true,
}

// @brief 上传图片
// @note 这里为了方便直接使用了阿里云OSS作为存储容器,当然了,也可以使用/admin/setting里的form自行配置
func UploadImg(c *gin.Context, picName string) (string, error) {
	var ossStatus = true
	if !ossStatus {
		return LocalUploadImg(c, picName)
	}
	return OssUploadImg(c, picName)
}

func OssUpload(file *multipart.FileHeader, dst string) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	err = utils.OSSBucket.PutObject(dst, f)
	if err != nil {
		return "", err
	}
	return dst, nil
}

// @brief 保存图片到阿里云OSS中
func OssUploadImg(c *gin.Context, picName string) (string, error) {
	file, err := c.FormFile(picName)

	if err != nil {
		return "", err
	}

	extName := path.Ext(file.Filename)
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀不合法")
	}

	day := models.GetDay()
	dir := UPLOAD_DIR + day

	fileName := strconv.FormatInt(models.TimeStampNano(), 10) + extName

	dst := path.Join(dir, fileName)
	// 调用oss上传:
	OssUpload(file, dst)

	return dst, nil
}

// @brief 上传图片保存到本地
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
	// 获取上传的文件:
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}

	// 获取后缀名判断是否是图片: .jpg, .png, .gif, .jpeg
	extName := path.Ext(file.Filename)
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}

	// 创建图片保存的目录
	day := models.GetDay()
	dir := UPLOAD_DIR + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// 生成文件名称和文件保存的目录
	fileName := models.Int2Str(int(models.TimeStampNano())) + extName

	// 上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}

// 格式化输出图片
func FormatImg(str string) string {
	var ossStatus = 1
	if ossStatus == 1 {
		return models.Loader.AliOOSConfig.Domain + str
	} else {
		return "/" + str
	}
}

// @brief 商品缩略图
