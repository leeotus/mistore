package admin

import (
	"fmt"
	"mistore/src/models"
	"mistore/src/utils"
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
	// fmt.Println(userInput)

	if !userInput.VerifyCaptcha() {
		ctl.Error(c, "验证码验证失败", "/admin/login")
		return
	}

	/*
		// @note: 原本的使用session来保存用户信息
		if res := userInput.VerifyCaptcha(); res {
			if ok, user := userInput.LogIn(); ok {
				// 设置用户的session
				sessionId := models.GenerateSessionUUID()

				// 传入上下文
				err := userInput.SaveSession(c.Request.Context(), sessionId, user.Username, user.IsSuper)
				if err != nil {
					ctl.Error(c, "会话创建失败", "/admin/login")
					return
				}

				// 将SessionID写入到Cookie,发送给客户端
				c.SetCookie(models.COOKIE_SESSNAME, sessionId, 3600, "/", "localhost", false, true)

				ctl.Success(c, "登录成功", "/admin")
			} else {
				ctl.Error(c, "用户名或密码错误", "/admin/login")
			}
		} else {
			ctl.Error(c, "验证码验证失败", "/admin/login")
		}
	*/
	// 验证用户登录
	if ok, user := userInput.LogIn(); ok {
		// 生成JWT令牌
		accessToken, err := utils.GenerateAccessToken(user)
		if err != nil {
			ctl.Error(c, "令牌生成失败", "/admin/login")
			return
		}
		// 生成刷新令牌
		refreshToken, err := utils.GenerateRefreshToken(user)
		if err != nil {
			ctl.Error(c, "刷新令牌生成失败", "/admin/login")
			return
		}

		// 设置JWT响应头
		c.Header("Authorization", "Bearer "+accessToken)
		c.Header("X-Refresh-Token", refreshToken)

		// 返回成功响应
		// res, _ := json.Marshal(gin.H{
		// 	"code": 200,
		// 	"msg":  "登录成功",
		// 	"data": gin.H{
		// 		"access_token":  accessToken,
		// 		"refresh_token": refreshToken,
		// 		"expires_in":    int64(models.AccessTokenDuration.Seconds()),
		// 	},
		// })
		// fmt.Println("登录成功:", string(res))
		ctl.Success(c, "登录成功", "/admin")
		// 返回JSON响应，包含重定向URL
		// c.JSON(http.StatusOK, gin.H{
		// 	"code":     http.StatusOK,
		// 	"msg":      "登录成功",
		// 	"redirect": "/admin",
		// })
	} else {
		ctl.Error(c, "用户名或密码错误", "/admin/login")
	}

}

/**
 * @brief 登出操作
 */
func (ctl LoginController) LoginOut(c *gin.Context) {
	// 获取当前用户信息
	claims, exists := c.Get("user_claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未登录"})
		return
	}

	// @todo 此处可以添加将token加入黑名单的逻辑

	// 清除客户端的认证信息
	c.Header("Authorization", "")
	c.Header("X-Refresh-Token", "")

	fmt.Println("登出成功", claims)
	c.Redirect(http.StatusFound, "/admin/login")
	/*
		sessionId, err := c.Cookie(models.COOKIE_SESSNAME)

		// 删除该用户的Session
		if err == nil && len(sessionId) != 0 {
			redisKey := fmt.Sprintf("%s%s", models.SESSION_PREFIX, sessionId)

			ctx := c.Request.Context()
			_, delErr := db.RedisDB.Del(ctx, redisKey).Result()
			if delErr != nil {
				// TODO: 销毁session失败的情况
			}
		}

		// 清除客户端的session_id Cookie
		// NOTE: 设置Cookie值为空表示删除
		c.SetCookie(models.COOKIE_SESSNAME, "", -1, "/", "localhost", false, true)

		// 重定向:
		c.Redirect(http.StatusFound, "/admin/login")
	*/
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

// @brief 刷新令牌
func (ctl LoginController) RefreshToken(c *gin.Context) {
	refreshToken := c.GetHeader("X-Refresh-Token")
	if refreshToken == "" {
		ctl.Error(c, "缺少刷新令牌", "/admin/login")
		return
	}

	// 验证并刷新令牌
	newTokens, err := utils.RefreshTokens(refreshToken)
	if err != nil {
		ctl.Error(c, "无效的刷新令牌", "/admin/login")
		return
	}

	// 设置新的token
	c.Header("Authorization", "Bearer "+newTokens.AccessToken)
	c.Header("X-Refresh-Token", newTokens.RefreshToken)

	// 返回新的令牌
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "令牌刷新成功",
		"data": newTokens,
	})
}
