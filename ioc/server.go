package ioc

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/internal/handler"
	"github.com/lcsin/pandora/web"
)

// InitWebServer 初始化web服务
//
//	@param userHandler
//	@return *gin.Engine
func InitWebServer(userHandler *handler.UserHandler,
	musicHandler *handler.MusicHandler, middlewares []gin.HandlerFunc) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares...)
	// 只是为了将asset下的静态资源在编译器期打包进可执行的二进制文件
	r.SetHTMLTemplate(template.Must(template.New("tmpl").ParseFS(web.StaticFS, "assets/*")))
	// 绑定静态资源请求（将静态资源的请求/assets/**，映射到./web/static目录下）
	r.Static("/assets", "./web/assets")
	r.StaticFS("/music", gin.Dir("music", true))
	

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 前端路由
	// WebHandler.RegisterRoutes(r)

	// 后端接口
	v1 := r.Group("/api/v1")
	userHandler.RegisterRoutes(v1)  // 注册用户hanlder
	musicHandler.RegisterRoutes(v1) // 注册音乐hanlder

	return r
}
