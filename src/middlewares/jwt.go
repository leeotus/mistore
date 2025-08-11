package middlewares

import (
	"mistore/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWT认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		// 检查是否需要fresh token
		if utils.RequireFreshToken(claims) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "fresh token required",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("is_admin", claims.IsAdmin)
		ctx.Next()
	}
}

// 刷新token
func FreshTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			ctx.Abort()
			return
		}

		// 检查是否为Fresh Token
		if !claims.IsFresh {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "fresh token required",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
