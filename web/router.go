package web

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler web handler
// type Handler struct {
// 	musicSrv service.IMusicService
// }

// NewHandler web handler构造器
//
//	@param musicSrv
//	@return *Handler
// func NewHandler(musicSrv service.IMusicService) *Handler {
// 	return &Handler{musicSrv: musicSrv}
// }

// RegisterRoutes 注册路由
//
//	@receiver wh
//	@param r
func RegisterRoutes(r gin.IRouter) {
	// r.GET("/login", Login)
	// r.GET("/register", Register)

	r.GET("/", Index)
	r.GET("/index", Index)
}

// Login 登录页面
//
//	@param c
// func Login(c *gin.Context) {
// 	c.HTML(http.StatusOK, "login.html", nil)
// }

// Register 注册页面
//
//	@param c
// func Register(c *gin.Context) {
// 	c.HTML(http.StatusOK, "register.html", nil)
// }

// Index 主页面
//
//	@param c
func Index(c *gin.Context) {
	tmpl := template.Must(template.New("tmpl").Parse(HTML))
	var buf bytes.Buffer
	tmpl.Execute(&buf, nil)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, buf.String())
}
