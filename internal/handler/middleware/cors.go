package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS 跨域中间件
//
//	@return gin.HandlerFunc
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		// 是否允许带上用户认证比如cookie
		AllowCredentials: true,
		// 业务请求中可以带上的请求头
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 哪些来源是允许的
		AllowOrigins: []string{"*"},
		// AllowOriginFunc: func(origin string) bool {
		// 	// 开发环境
		// 	if strings.HasPrefix(origin, "http://localhost") {
		// 		return true
		// 	}
		// 	// 生产环境
		// 	return strings.Contains(origin, "your_company.com")
		// },
		MaxAge: 12 * time.Hour,
	})
}
