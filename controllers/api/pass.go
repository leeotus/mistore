package api

import (
	"fmt"
	"mistore/src/models"
	"mistore/src/verify"
	"net/http"

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
	// 生成随机数
	fmt.Println(models.GetRandomNum())
	c.String(http.StatusOK, "login")
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
