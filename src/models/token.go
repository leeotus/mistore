package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenDuration  = time.Hour * 2
	RefreshTokenDuration = time.Hour * 24
	FreshTokenDuration   = time.Minute * 15
)

// @brief JWT Claims
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	IsFresh  bool   `json:"is_fresh"`
	jwt.RegisteredClaims
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}
