package main

import (
	"mistore/routers"
	"mistore/src/db"
	"mistore/src/models"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	loader := &models.IniLoader{}
	loader.LoadConfigs("./config/mistore.ini")
	db.MysqlCoreInit(loader.MysqlConfig.Username,
		loader.MysqlConfig.Password,
		loader.MysqlConfig.HostAddr,
		loader.MysqlConfig.Port,
		loader.MysqlConfig.DBName)

	db.RedisCoreInit(loader.RedisConfig.Host,
		loader.RedisConfig.Port,
		loader.RedisConfig.Password,
		0, 12, 2, 300*time.Second, 60*time.Second)

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"TimeStamp":  models.TimeStamp,
		"UnixToTime": models.UnixToTime,
		"Md5":        models.Md5,
		"UUID":       models.GenerateSessionUUID,
	})

	r.Static("admin/bootstrap/", "static/admin/bootstrap/")
	r.Static("admin/css/", "static/admin/css")
	r.Static("admin/js/", "static/admin/js")
	r.Static("admin/images/", "static/admin/images")

	r.LoadHTMLGlob("templates/**/**/*")

	r.GET("/favicon.ico", func(ctx *gin.Context) {
		data, err := os.ReadFile("static/favicon.ico")
		if err != nil {
			ctx.AbortWithStatus(500)
			os.Exit(1)
		}
		ctx.Data(http.StatusOK, "image/png", data)
	})

	routers.AdminRouterInit(r)

	r.Run()
}
