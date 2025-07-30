package admin

import (
	"mistore/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @brief 轮播图实现
 * @author leeotus (leeotus@163.com)
 */

type FocusController struct {
	controllers.BaseController
}

func (ctl FocusController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/index.html", gin.H{})
}

func (ctl FocusController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (ctl FocusController) Edit(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}

func (ctl FocusController) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "delete ok")
}
