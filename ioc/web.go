package ioc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/internal/handler"
)

// InitWebServer 初始化web服务
//
//	@param userHandler
//	@return *gin.Engine
func InitWebServer(webHandler *handler.WebHandler, userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 注册页面hanlder
	webHandler.RegisterRoutes(r)

	// 接口版本
	v1 := r.Group("/api/v1")
	// 注册用户hanlder
	userHandler.RegisterRoutes(v1)

	return r
}
