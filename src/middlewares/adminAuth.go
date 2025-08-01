package middlewares

import (
	"mistore/src/models"
	"net/http"
	"strings"
	"time"

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
	sessionId, err := ctx.Cookie(models.COOKIE_SESSNAME)
	if err != nil || sessionId == "" {
		// 没有找到session_id,需要重定向到登录页
		ctx.Redirect(http.StatusFound, "/admin/login")
		ctx.Abort() // 不执行后续的中间件函数
		return
	}

	// 用新Session类获取session数据
	sessKey := models.SESSION_PREFIX + sessionId
	sessData, err := models.GetSession(ctx, sessKey)
	if err != nil || sessData == nil {
		// Session不存在或者已经过期,需要重定向到登录页面
		ctx.Redirect(http.StatusFound, "/admin/login")
		ctx.Abort()
		return
	}

	// 滑动过期：每次操作距离上次刷新超过3分钟才刷新session过期时间
	models.RefreshSessionWithInterval(ctx, sessKey, sessData, 3*time.Minute)

	isSuper, _ := sessData["super_user"].(float64) // json解码数字为float64
	if int(isSuper) != 1 {
		ctx.String(http.StatusForbidden, "请联系管理员获取权限!")
		ctx.Abort()
		return
	}
	ctx.Next()
}
