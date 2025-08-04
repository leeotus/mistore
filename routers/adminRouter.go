package routers

import (
	"mistore/controllers/admin"
	"mistore/src/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	admin_router := r.Group("/admin", middlewares.InitAdminAuthMiddleware)

	{
		admin_router.GET("/", admin.MainPageController{}.Index)
		admin_router.GET("/welcome", admin.MainPageController{}.Welcome)
		admin_router.GET("/changestatus", admin.MainPageController{}.ChangeStatus)

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
		admin_router.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)

		// 轮播图
		admin_router.GET("/focus", admin.FocusController{}.Index)
		admin_router.GET("/focus/add", admin.FocusController{}.Add)
		admin_router.GET("/focus/edit", admin.FocusController{}.Edit)
		admin_router.GET("/focus/delete", admin.FocusController{}.Delete)
		admin_router.POST("/focus/doAdd", admin.FocusController{}.DoAdd)
		admin_router.POST("/focus/doEdit", admin.FocusController{}.DoEdit)

		// 权限视图
		admin_router.GET("/role", admin.RoleController{}.Index)
		admin_router.GET("/role/add", admin.RoleController{}.Add)
		admin_router.GET("/role/edit", admin.RoleController{}.Edit)
		admin_router.GET("/role/delete", admin.RoleController{}.Delete)
		admin_router.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		admin_router.POST("/role/doEdit", admin.RoleController{}.DoEdit)

		// 商品分类页面
		admin_router.GET("/goodsCate", admin.GoodsCateController{}.Index)
		admin_router.GET("/goodsCate/add", admin.GoodsCateController{}.Add)
		admin_router.GET("/goodsCate/edit", admin.GoodsCateController{}.Edit)
		admin_router.GET("/goodsCate/delete", admin.GoodsCateController{}.Delete)
		admin_router.POST("/goodsCate/doAdd", admin.GoodsCateController{}.DoAdd)
		admin_router.POST("/goodsCate/doEdit", admin.GoodsCateController{}.DoEdit)

		// 商品类型
		admin_router.GET("/goodsType", admin.GoodsTypeController{}.Index)
		admin_router.GET("/goodsType/add", admin.GoodsTypeController{}.Add)
		admin_router.GET("/goodsType/edit", admin.GoodsTypeController{}.Edit)
		admin_router.GET("/goodsType/delete", admin.GoodsTypeController{}.Delete)
		admin_router.POST("/goodsType/doAdd", admin.GoodsTypeController{}.DoAdd)
		admin_router.POST("/goodsType/doEdit", admin.GoodsTypeController{}.DoEdit)

		// 商品属性
		admin_router.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController{}.Index)
		admin_router.GET("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.Add)
		admin_router.GET("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeController{}.Edit)
		admin_router.POST("/goodsTypeAttribute/doAdd", admin.GoodsTypeAttributeController{}.DoAdd)
		admin_router.POST("/goodsTypeAttribute/doEdit", admin.GoodsTypeAttributeController{}.DoEdit)

		// 商品页面
		admin_router.GET("/goods", admin.GoodsController{}.Index)
		admin_router.GET("/goods/add", admin.GoodsController{}.Add)
		admin_router.GET("/goods/changeGoodsImageColor", admin.GoodsController{}.ChangeGoodsImageColor)
		admin_router.GET("/goods/removeGoodsImage", admin.GoodsController{}.RemoveGoodsImage)
		admin_router.GET("/goods/goodsTypeAttribute", admin.GoodsController{}.GoodsTypeAttribute)
		admin_router.GET("/goods/edit", admin.GoodsController{}.Edit)
		admin_router.GET("/goods/delete", admin.GoodsController{}.Delete)
		admin_router.POST("/goods/doAdd", admin.GoodsController{}.DoAdd)
		admin_router.POST("/goods/imageUpload", admin.GoodsController{}.ImageUpload)
		admin_router.POST("/goods/doEdit", admin.GoodsController{}.DoEdit)
	}
}
