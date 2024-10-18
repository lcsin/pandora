package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/api"
	"github.com/lcsin/pandora/api/message"
	"github.com/lcsin/pandora/internal/service"
)

// MusicHandler music handler
type MusicHandler struct {
	musicSrv service.IMusicService
}

// NewMusicHandler music handler 构造函数
//
//	@param musicSrv
//	@return *MusicHandler
func NewMusicHandler(musicSrv service.IMusicService) *MusicHandler {
	return &MusicHandler{musicSrv: musicSrv}
}

// RegisterRoutes 注册音乐服务路由
//
//	@receiver mh
//	@param v1
func (mh *MusicHandler) RegisterRoutes(v1 *gin.RouterGroup) {
	mg := v1.Group("/music")

	mg.POST("/list", mh.GetMusicListByUID)
}

// GetMusicListByUID 根据用户id获取音乐列表
//
//	@receiver mh
//	@param c
func (mh *MusicHandler) GetMusicListByUID(c *gin.Context) {
	type Req struct {
		UID int64 `json:"uid"`
	}

	var req Req
	if err := c.ShouldBind(&req); err != nil {
		api.ResponseError(c, message.BadRequest)
		return
	}
	if req.UID == 0 {
		api.ResponseErrorMessage(c, message.BadRequest, "用户ID不能为空")
		return
	}

	music, err := mh.musicSrv.GetMusicListByUID(c, req.UID)
	if err != nil {
		api.ResponseError(c, message.Failed)
		return
	}

	api.ResponseOK(c, music)
}
