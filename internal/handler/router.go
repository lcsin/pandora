package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/api"
	"github.com/lcsin/pandora/api/message"
	"github.com/lcsin/pandora/internal/service"
)

// WebHandler web handler
type WebHandler struct {
	musicSrv service.IMusicService
}

// NewWebHandler web handler构造器
//
//	@param musicSrv
//	@return *Handler
func NewWebHandler(musicSrv service.IMusicService) *WebHandler {
	return &WebHandler{musicSrv: musicSrv}
}

// RegisterRoutes 注册路由
//
//	@receiver wh
//	@param r
func (wh *WebHandler) RegisterRoutes(r gin.IRouter) {
	r.GET("/login", wh.Login)
	r.GET("/register", wh.Register)
	r.GET("/index", wh.Index)
}

// Login 登录页面
//
//	@param c
func (wh *WebHandler) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"email":    "1847@qq.com",
		"password": "root",
	})
}

// Register 注册页面
//
//	@param c
func (wh *WebHandler) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// Index 主页面
//
//	@param c
func (wh *WebHandler) Index(c *gin.Context) {
	type Music struct {
		Name   string `json:"name"`
		Author string `json:"author"`
		Time   string `json:"time"`
		URL    string `json:"url"`
	}

	// music := []Music{
	// 	{Name: "あとひとつ", Author: "FUNKY MONKEY BABYS (放克猴宝贝)", Time: "5.10", URL: "../static/test1.ogg"},
	// 	{Name: "三線の花", Author: "BEGIN", Time: "5.10", URL: "../static/test2.mp3"},
	// 	{Name: "サイレントマジョリティー", Author: "欅坂46", Time: "5.10", URL: "../static/test3.mp3"},
	// }
	// fmt.Println(music)

	resp, err := wh.musicSrv.GetMusicListByUID(c, 2)
	if err != nil {
		api.ResponseError(c, message.Failed)
		return
	}
	music := make([]Music, 0, len(resp))
	for _, v := range resp {
		music = append(music, Music{
			Name:   v.Name,
			Author: v.Author,
			Time:   v.Time,
			URL:    v.URL,
		})
	}

	c.HTML(http.StatusOK, "index.html", music)
}
