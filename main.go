/**
 * @brief gin商场实战项目
 * @author leeotus
 * @email leeotus@163.com
 */
package main

import (
	"mistore/controllers/admin"
	"mistore/routers"
	"mistore/src/db"
	"mistore/src/models"
	"mistore/src/utils"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// NOTE: 其实也可以将整个main里的主要代码进行封装,写入到另一个go文件里,只在main里暴露启动服务器的代码
func main() {
	models.LoadConfigs("./config/mistore.ini") // NOTE: 需要修改为自己的ini文件,具体看config/example.ini

	// @brief 初始化MySQL数据库
	db.MysqlCoreInit(models.Loader.MysqlConfig.Username,
		models.Loader.MysqlConfig.Password,
		models.Loader.MysqlConfig.HostAddr,
		models.Loader.MysqlConfig.Port,
		models.Loader.MysqlConfig.DBName)

	// @brief 初始化Redis数据库
	db.RedisCoreInit(models.Loader.RedisConfig.Host,
		models.Loader.RedisConfig.Port,
		models.Loader.RedisConfig.Password,
		0, 12, 2, 300*time.Second, 60*time.Second)

	// @brief 初始化阿里云OSS
	utils.OOSClientInitFromIni(models.Loader.AliOOSConfig.EndPoint, models.Loader.AliOOSConfig.AccessKey, models.Loader.AliOOSConfig.AccessSecret)

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"TimeStamp":  models.TimeStamp,
		"UnixToTime": models.UnixToTime,
		"Md5":        models.Md5,
		"UUID":       models.GenerateSessionUUID,
		"Str2Html":   models.Str2Html,
		"FormatImg":  admin.FormatImg,
		"Sub":        models.Sub,
		"Substr":     models.Substr,
	})

	r.Static("admin/bootstrap/", "static/admin/bootstrap/")
	r.Static("admin/css/", "static/admin/css")
	r.Static("admin/js/", "static/admin/js")
	r.Static("admin/images/", "static/admin/images")
	r.Static("static/", "static/")

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
	routers.DefaultRoutersInit(r)

	r.Run()
}
