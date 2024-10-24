package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT载荷
type Claims struct {
	jwt.RegisteredClaims
	UID int64 `json:"uid"`
	// UserAgent string `json:"user_agent"`
}

// TokenInfo Token信息
type TokenInfo struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenExpire int64  `json:"accessTokenExpire"`
}
