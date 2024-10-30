package web

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
//
//	@receiver wh
//	@param r
func RegisterRoutes(r gin.IRouter) {
	r.GET("/", Index)
	r.GET("/index", Index)
}

// Index 主页面
//
//	@param c
func Index(c *gin.Context) {
	tmpl := template.Must(template.New("tmpl").Parse(IndexHTML))
	var buf bytes.Buffer
	tmpl.Execute(&buf, nil)

	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, buf.String())
}
