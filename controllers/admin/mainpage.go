package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainPageController struct{}

func (ctl MainPageController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/mainpage/index.html", gin.H{})
}

func (ctl MainPageController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/mainpage/welcome.html", gin.H{})
}
