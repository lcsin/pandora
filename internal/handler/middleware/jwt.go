package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lcsin/pandora/api"
	"github.com/lcsin/pandora/api/message"
	"github.com/lcsin/pandora/internal/domain"
	"github.com/spf13/viper"
)

// JWTBuilder jwt builder
type JWTBuilder struct {
	ignores map[string]bool
	statics map[string]struct{}
}

// NewJWTBuilder new jwt builder
//
//	@return *JWTBuilder
func NewJWTBuilder() *JWTBuilder {
	return &JWTBuilder{}
}

// Ignores 忽略HTTP请求
//
//	@receiver j
//	@param paths
//	@return *JWTBuilder
func (j *JWTBuilder) Ignores(paths ...string) *JWTBuilder {
	j.ignores = make(map[string]bool, len(paths))
	for _, v := range paths {
		j.ignores[v] = true
	}

	return j
}

// Statics 忽略静态资源请求
//
//	@receiver j
//	@param paths
//	@return *JWTBuilder
func (j *JWTBuilder) Statics(paths ...string) *JWTBuilder {
	j.statics = make(map[string]struct{}, len(paths))
	for _, v := range paths {
		j.statics[v] = struct{}{}
	}

	return j
}

// Build build jwt middleware
//
//	@receiver j
//	@return gin.HandlerFunc
func (j *JWTBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行需要忽略的请求
		path := c.Request.URL.Path
		if j.ignores[path] {
			return
		}
		// 放行静态资源
		for k := range j.statics {
			if strings.HasPrefix(path, k) {
				return
			}
		}

		// 获取请求头中的token
		headerToken := c.GetHeader("Authorization")
		if headerToken == "" {
			api.ResponseError(c, message.Unauthorized)
			c.Abort()
			return
		}

		// 解析token
		var claims domain.Claims
		token, err := jwt.ParseWithClaims(headerToken, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_secret")), nil
		})
		if err != nil {
			api.ResponseError(c, message.Unauthorized)
			c.Abort()
			return
		}

		// 校验token
		if !token.Valid {
			api.ResponseError(c, message.Unauthorized)
			c.Abort()
			return
		}

		// gin.Context上下文添加UID
		c.Set("uid", claims.UID)
	}
}
