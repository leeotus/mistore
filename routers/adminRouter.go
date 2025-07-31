package routers

import (
	"mistore/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	admin_router := r.Group("/admin")

	{
		admin_router.GET("/", admin.MainPageController{}.Index)
		admin_router.GET("/welcome", admin.MainPageController{}.Welcome)

		// 登录界面
		admin_router.GET("/login", admin.LoginController{}.Index)
		admin_router.GET("/code", admin.LoginController{}.GenerateCode)
		admin_router.POST("/doLogin", admin.LoginController{}.DoLogin)
		admin_router.GET("/loginOut", admin.LoginController{}.LoginOut)

		// 管理员界面
		admin_router.GET("/manager", admin.ManagerController{}.Index)
		admin_router.GET("/manager/add", admin.ManagerController{}.Add)
		admin_router.GET("/manager/edit", admin.ManagerController{}.Edit)
		admin_router.GET("/manager/delete", admin.ManagerController{}.Delete)

		// 轮播图
		admin_router.GET("/focus", admin.FocusController{}.Index)
		admin_router.GET("/focus/add", admin.FocusController{}.Add)
		admin_router.GET("/focus/edit", admin.FocusController{}.Edit)
		admin_router.GET("/focus/delete", admin.FocusController{}.Delete)
	}
}
