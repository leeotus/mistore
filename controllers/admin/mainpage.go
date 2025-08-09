package admin

import (
	"mistore/src/db"
	"mistore/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainPageController struct {
	BaseController
}

// @brief 后台管理主界面
func (ctl MainPageController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/mainpage/index.html", gin.H{})
}

// @brief 欢迎界面
func (ctl MainPageController) Welcome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/mainpage/welcome.html", gin.H{})
}

// @brief 公共修改状态的方法
func (ctl MainPageController) ChangeStatus(ctx *gin.Context) {
	id, err1 := models.Str2Int(ctx.Query("id"))
	if err1 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "传入的参数有误的",
		})
		return
	}

	table := ctx.Query("table")
	field := ctx.Query("field")

	// err := db.MySQLDB.Exec("update ? set ?=ABS(?-1) where id=?", table, field, field, id).Error
	err := db.MySQLDB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败,请重新尝试",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

func (ctl MainPageController) FlushAll(ctx *gin.Context) {
	db.RedisDB.FlushAll(ctx.Request.Context())
	ctl.Success(ctx, "清理Redis缓存成功", "/admin")
}
