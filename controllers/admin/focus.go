package admin

import (
	"errors"
	"fmt"
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 轮播图实现
 * @author leeotus (leeotus@163.com)
 */

type FocusController struct {
	BaseController
}

const UPLOAD_DIR = "./static/upload/"

var allowExtMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".jpeg": true,
	".gif":  true,
}

func UploadImg(c *gin.Context, picName string) (string, error) {
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
	fileName := models.Int2Str(int(models.TimeStamp())) + extName

	// 上传
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}

func (ctl FocusController) Index(ctx *gin.Context) {
	focusList := []models.Focus{}
	db.MySQLDB.Find(&focusList)
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (ctl FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (ctl FocusController) DoAdd(ctx *gin.Context) {
	title := strings.TrimSpace(ctx.PostForm("title"))
	focusType, _ := models.Str2Int(ctx.PostForm("focus_type"))
	link := ctx.PostForm("link")
	sort, err := models.Str2Int(ctx.PostForm("sort"))
	if err != nil {
		ctl.Error(ctx, "排序输入非法!", "admin/focus/add")
		return
	}
	status, _ := models.Str2Int(ctx.PostForm("status"))

	// 使用封装好的UploadImg来上传文件, 把文件上传的任务放到线程池中
	focusImg, err2 := UploadImg(ctx, "focus_img")
	if err2 != nil {
		fmt.Println(err2)
	}

	focus := models.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImg,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(models.TimeStamp()),
	}
	err = db.MySQLDB.Create(&focus).Error
	if err != nil {
		ctl.Error(ctx, "增加轮播图失败", "/admin/focus/add")
	} else {
		ctl.Success(ctx, "增加轮播图成功", "/admin/focus")
	}
}

func (ctl FocusController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}

func (ctl FocusController) DoEdit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "do edit")
}

func (ctl FocusController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete ok")
}
