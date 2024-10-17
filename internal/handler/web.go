package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WebHandler web page handler
type WebHandler struct {
}

// NewWebHandler web handler 构造函数
//
//	@return *WebHandler
func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

// RegisterRoutes 注册路由
//
//	@receiver wh
//	@param r
func (wh *WebHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/login", wh.Login)
	r.GET("/register", wh.Register)
}

// Login 登录页面
//
//	@receiver sh
//	@param c
func (wh *WebHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"email":    "1847@qq.com",
		"password": "root",
	})
}

// Register 注册页面
//
//	@receiver wh
//	@param c
func (wh *WebHandler) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
