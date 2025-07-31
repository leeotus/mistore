package middlewares

import (
	"mistore/src/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @note 不需要认证的路径
var NoAuthPaths = []string{
	"/admin/login",
	"/admin/css/",
	"/admin/js/",
	"/admin/images/",
	"/admin/code",
	"/admin/doLogin",
}

func InitAdminAuthMiddleware(ctx *gin.Context) {
	// 获取用户访问的url
	pathname := strings.Split(ctx.Request.URL.String(), "?")[0]

	for _, path := range NoAuthPaths {
		if strings.HasPrefix(pathname, path) {
			ctx.Next() // 直接运行下一个中间件函数
			return
		}
	}

	// RESEARCH: 将SessionID保存在客户端的Cookie中,如何才能最大限度的保证安全呢?
	// 获取用户的Session信息
	sessionId, err := ctx.Cookie("session_id")
	if err != nil || sessionId == "" {
		// 没有找到session_id,需要重定向到登录页
		ctx.Redirect(http.StatusFound, "/admin/login")
		ctx.Abort() // 不执行后续的中间件函数
		return
	}

	// 从redis中获取用户的信息
	redisKey := "session:" + sessionId
	userId, err := db.RedisDB.HGet(ctx, redisKey, "session_val").Result()
	if err != nil || userId == "" {
		// Session不存在或者已经过期,需要重定向到登录页面
		ctx.Redirect(http.StatusFound, "/admin/login")
		ctx.Abort()
		return
	}

	// 验证是否为管理员身份,admin界面只向管理员开放
	isSuper, _ := db.RedisDB.HGet(ctx, redisKey, "super_user").Result()
	isAdmin := checkIsAdmin(isSuper)
	if !isAdmin {
		ctx.String(http.StatusForbidden, "请联系管理员获取权限!")
		ctx.Abort()
		return
	}
	ctx.Next()
}

func checkIsAdmin(issuper string) bool {
	return issuper == "1"
}
