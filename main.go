package main

import (
	"fmt"
	"mistore/routers"
	"net/http"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{})

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
