package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 管理员登录完成后的界面
 * @author leeotus (leeotus@163.com)
 */

type ManagerController struct {
	BaseController
}

// TODO: 管理员界面下的各种操作

func (ctl ManagerController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/index.html", gin.H{})
}

func (ctl ManagerController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/add.html", gin.H{})
}

func (ctl ManagerController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{})
}

func (ctl ManagerController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete ok")
}
