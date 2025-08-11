package utils

import (
	"errors"
	"mistore/src/db"
	"mistore/src/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// @todo secret-key
var jwtSecret = []byte("your-secret-key")

// @brief 生成JWT访问令牌
func GenerateAccessToken(user any) (string, error) {
	var claims *models.JWTClaims
	switch u := user.(type) {
	// 之后可以扩展其他的类型
	case models.Manager:
		claims = &models.JWTClaims{
			UserID:   uint(u.Id),
			Username: u.Username,
			IsAdmin:  true,
			IsFresh:  true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(models.AccessTokenDuration)), // 令牌过期时间
				IssuedAt:  jwt.NewNumericDate(time.Now()),                                 // 令牌签发时间
				Issuer:    "mistore",                                                      // 表示令牌的签发者
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// @brief 生成刷新令牌
func GenerateRefreshToken(user any) (string, error) {
	var claims *models.JWTClaims
	switch u := user.(type) {
	// 之后可以扩展其他的类型
	case models.Manager:
		claims = &models.JWTClaims{
			UserID:   uint(u.Id),
			Username: u.Username,
			IsAdmin:  true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(models.RefreshTokenDuration)), // 令牌过期时间
				IssuedAt:  jwt.NewNumericDate(time.Now()),                                  // 令牌签发时间
				Issuer:    "mistore",                                                       // 表示令牌的签发者
			},
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// @brief 验证令牌
func ValidateToken(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(t *jwt.Token) (any, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// @brief 刷新令对
func RefreshTokens(refreshToken string) (*models.TokenResponse, error) {
	claims, err := ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	// 获取用户信息
	if claims.IsAdmin {
		// 说明是管理员用户
		var user = models.Manager{
			Id: int(claims.UserID),
		}
		res := db.MySQLDB.Find(&user)
		if res.RowsAffected == 0 {
			// 说明没有找到用户
			return nil, errors.New("user not found")
		}

		// 生成新的令对
		accessToken, err1 := GenerateAccessToken(user)
		if err1 != nil {
			return nil, err1
		}

		newRefreshToken, err2 := GenerateRefreshToken(user)
		if err2 != nil {
			return nil, err2
		}

		return &models.TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: newRefreshToken,
			ExpiresIn:    int64(models.AccessTokenDuration.Seconds()),
			TokenType:    "Bearer",
		}, nil
	}
	// @todo 其他类型
	return nil, errors.New("user type error")
}

// @brief 检查是否需要刷新token
func RequireFreshToken(claims *models.JWTClaims) bool {
	return !claims.IsFresh || claims.ExpiresAt.Before(time.Now().Add(-models.FreshTokenDuration))
}
