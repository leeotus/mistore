package admin

import (
	"fmt"
	"mistore/src/db"
	"mistore/src/models"
	"mistore/src/verify"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 管理员登录界面的controller
 * @author leeotus (leeotus@163.com)
 */

type LoginController struct {
	BaseController
}

/**
 * @brief 显示登录主页面的方法
 * @param *gin.Context
 */
func (ctl LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

/**
 * @brief 登录操作
 * @note 登录界面填写用户,密码后提交表单执行登录
 */
func (ctl LoginController) DoLogin(c *gin.Context) {
	// 需要判断输入的验证码是否正确
	userInput := &models.LoginData{
		CapId:     c.PostForm("captchaId"),
		VerifyVal: c.PostForm("verifyValue"),
		Username:  c.PostForm("username"),
		Password:  c.PostForm("password"),
	}
	fmt.Println(userInput)

	if res := userInput.VerifyCaptcha(); res {
		if ok, user := userInput.LogIn(); ok {
			// 设置用户的session
			sessionId := models.GenerateSessionUUID()

			// 传入上下文
			err := userInput.SaveSession(c.Request.Context(), sessionId, user.Username, user.IsSuper)
			if err != nil {
				ctl.error(c, "会话创建失败", "/admin/login")
				return
			}

			// 将SessionID写入到Cookie,发送给客户端
			fmt.Println("SessionID:", sessionId)
			c.SetCookie("session_id", sessionId, 3600, "/", "localhost", false, true)

			ctl.success(c, "登录成功", "/admin")
		} else {
			ctl.error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		ctl.error(c, "验证码验证失败", "/admin/login")
	}
}

/**
 * @brief 登出操作
 */
func (ctl LoginController) LoginOut(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")

	// 删除该用户的Session
	if err == nil && len(sessionId) != 0 {
		redisKey := fmt.Sprintf("session:%s", sessionId)

		ctx := c.Request.Context()
		_, delErr := db.RedisDB.Del(ctx, redisKey).Result()
		if delErr != nil {
			// TODO: 销毁session失败的情况
		}
	}

	// 清除客户端的session_id Cookie
	// NOTE: 设置Cookie值为空表示删除
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)

	// 重定向:
	c.Redirect(http.StatusFound, "/admin/login")
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
