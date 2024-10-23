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
func InitWebServer(WebHandler *handler.WebHandler, userHandler *handler.UserHandler, musicHandler *handler.MusicHandler) *gin.Engine {
	r := gin.Default()
	// 绑定HTML模板文件
	r.SetHTMLTemplate(template.Must(template.New("tmpl").ParseFS(web.HTMLFS, "template/*")))
	// 绑定静态资源文件（将静态资源请求/static/**，映射到./web/static目录下）
	r.Static("/static", "./web/static")
	r.StaticFS("/file", gin.Dir("file", true))

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 前端路由
	WebHandler.RegisterRoutes(r)

	// 后端接口
	v1 := r.Group("/api/v1")
	userHandler.RegisterRoutes(v1)  // 注册用户hanlder
	musicHandler.RegisterRoutes(v1) // 注册音乐hanlder

	return r
}
