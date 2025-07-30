package admin

import (
	"fmt"
	"mistore/controllers"
	"mistore/src/verify"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 管理员登录界面的controller
 * @author leeotus (leeotus@163.com)
 */

type LoginController struct {
	controllers.BaseController
}

/**
 * @brief 显示登录主页面的方法
 * @param *gin.Context
 */
func (ctl LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (ctl LoginController) DoLogin(c *gin.Context) {
	// 需要判断输入的验证码是否正确
	capId := c.PostForm("captchaId")
	verifyVal := c.PostForm("verifyValue")
	if res := verify.VerifyCaptcha(capId, verifyVal); res {
		// 验证成功
		// TODO: 需要把用户数据写入到admin的index页面
		// NOTE: 现在暂时先返回一个简单的string
		c.String(http.StatusOK, "验证码验证成功")
	} else {
		c.String(http.StatusOK, "验证码验证失败")
	}
}

// 登录界面处验证码相关的方法:
func (ctl LoginController) GenerateCode(c *gin.Context) {
	id, b64s, ans, err := verify.GenerateCaptcha()
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
		"captchaAns":   ans,
	})
}
