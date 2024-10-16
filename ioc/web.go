package ioc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/internal/handler"
)

// InitWebServer 初始化web服务
//
//  @param userHandler 
//  @return *gin.Engine 
func InitWebServer(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
