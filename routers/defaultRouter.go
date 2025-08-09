package routers

import (
	"mistore/controllers/api"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	default_router := r.Group("/")
	{
		default_router.GET("/", api.DefaultController{}.Index)
		default_router.GET("/category:id", api.ProductController{}.Category)
		default_router.GET("/detail", api.ProductController{}.Detail)
		default_router.GET("/product/getImgList", api.ProductController{}.GetImgList)

		default_router.GET("/pass/login", api.PassController{}.Login)
		default_router.GET("/pass/captcha", api.PassController{}.Captcha)
		default_router.GET("/pass/registerStep1", api.PassController{}.RegisterStep1)
		default_router.GET("/pass/registerStep2", api.PassController{}.RegisterStep2)
		default_router.GET("/pass/registerStep3", api.PassController{}.RegisterStep3)

		default_router.GET("/cart", api.CartController{}.Get)
		default_router.GET("/cart/addCart", api.CartController{}.AddCart)
		default_router.GET("/cart/successTip", api.CartController{}.AddCartSuccess)
		default_router.GET("/cart/incCart", api.CartController{}.IncCart)
		default_router.GET("/cart/decCart", api.CartController{}.DecCart)
		default_router.GET("/cart/changeOneCart", api.CartController{}.ChangeOneCart)
		default_router.GET("/cart/changeAllCart", api.CartController{}.ChangeAllCart)
		default_router.GET("/cart/delCart", api.CartController{}.DelCart)
	}
}
