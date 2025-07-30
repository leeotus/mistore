package admin

import (
	"mistore/controllers"
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

// TODO: 登录界面登录动作
func (ctl LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "登录中")
}
