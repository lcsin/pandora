package ioc

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/internal/handler/middleware"
)

// InitMiddlewares 初始化中间件
//
//	@return []gin.HandlerFunc
func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// JWT中间件
		middleware.NewJWTBuilder().
			Statics("/static", "/file"). // 放行静态资源
			Ignores(
				"/login", "/register", "/index", // 放行页面路由
				"/ping", "/api/v1/users/login", "/api/v1/users/register", // 放行后端请求
			).Build(),
		// 跨域中间件
		middleware.CORS(),
	}
}
