package api

import (
	"fmt"
	"mistore/src/db"
	"mistore/src/models"
	"mistore/src/utils"
	"mistore/src/verify"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type PassController struct {
	BaseController
}

func (con PassController) Captcha(c *gin.Context) {
	id, b64s, err := verify.GenerateCaptchaWithInputSize(50, 120, 4)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con PassController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "api/pass/login.html", gin.H{})
}

func (con PassController) RegisterStep1(c *gin.Context) {
	c.HTML(http.StatusOK, "api/pass/register_step1.html", gin.H{})
}

func (con PassController) RegisterStep2(c *gin.Context) {
	c.HTML(http.StatusOK, "api/pass/register_step2.html", gin.H{})
}

func (con PassController) RegisterStep3(c *gin.Context) {
	c.HTML(http.StatusOK, "api/pass/register_step3.html", gin.H{})
}

func (con PassController) DoLogin(c *gin.Context) {
	phone := strings.TrimSpace(c.PostForm("phone"))
	password := c.PostForm("password")
	captchaId := c.PostForm("captchaId")
	captchaVal := c.PostForm("captchaVal")

	// 1.验证图像验证码是否正确
	if flag := verify.VerifyCaptcha(captchaId, captchaVal); !flag {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "图像验证码不正确",
		})
		return
	}

	// 2.验证用户名密码是否正确
	password = models.Md5(strings.TrimSpace(password))
	userList := []models.User{}
	db.MySQLDB.Where("phone=? AND password=?", phone, password).Find(&userList)
	if len(userList) > 0 {
		fmt.Println("已找到用户,登录中...", userList[0])
		// @todo 改用JWT+fresh token来替换cookie
		utils.Cookie.Set(c, "userinfo", userList[0])
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "用户登录成功",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "用户名或者密码错误",
		})
		return
	}
}

func (con PassController) LoginOut(c *gin.Context) {
	// 删除Cookie里面的userinfo执行跳转
	// @todo jwt+fresh token
	utils.Cookie.Remove(c, "userinfo")
	prevPage := c.Request.Referer()
	if len(prevPage) > 0 {
		c.Redirect(http.StatusFound, prevPage)
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}
