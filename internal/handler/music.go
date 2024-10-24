package handler

import (
	"errors"
	"mime/multipart"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/api"
	"github.com/lcsin/pandora/api/message"
	"github.com/lcsin/pandora/internal/domain"
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
	mg.POST("/search", mh.GetMusicListByNameOrAuthor)
	mg.POST("/upload", mh.UploadMusic)
	mg.POST("/update", mh.UpdateMusic)
}

// GetMusicListByUID 根据用户id获取音乐列表
//
//	@receiver mh
//	@param c
func (mh *MusicHandler) GetMusicListByUID(c *gin.Context) {
	uid := c.GetInt64("uid")
	if uid == 0 {
		api.ResponseError(c, message.Unauthorized)
		return
	}

	music, err := mh.musicSrv.GetMusicListByUID(c, uid)
	if err != nil {
		api.ResponseError(c, message.Failed)
		return
	}

	api.ResponseOK(c, music)
}

// GetMusicListByNameOrAuthor 根据歌名或作者获取音乐列表
//
//	@receiver mh
//	@param c
func (mh *MusicHandler) GetMusicListByNameOrAuthor(c *gin.Context) {
	type Req struct {
		Query string `json:"query"`
	}

	var req Req
	if err := c.ShouldBind(&req); err != nil {
		api.ResponseError(c, message.BadRequest)
		return
	}

	music, err := mh.musicSrv.GetMusicListByNameOrAuthor(c, req.Query, req.Query)
	if err != nil {
		api.ResponseError(c, message.Failed)
		return
	}

	api.ResponseOK(c, music)
}

// UploadMusic 上传音乐
//
//	@receiver mh
//	@param c
func (mh *MusicHandler) UploadMusic(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		api.ResponseError(c, message.BadRequest)
		return
	}

	files := make([]*multipart.FileHeader, 0, len(form.File))
	for _, v := range form.File {
		// 校验文件类型是否为音频文件
		if !strings.HasPrefix(v[0].Header.Get("Content-Type"), "audio") {
			api.ResponseError(c, message.BadRequest)
			return
		}
		files = append(files, v...)
	}

	musics := make([]*domain.Music, 0, len(files))
	for _, f := range files {
		var name, author string
		names := strings.Split(strings.TrimSuffix(f.Filename, path.Ext(f.Filename)), "-")
		if len(names) == 2 {
			name = strings.TrimSpace(names[0])
			author = strings.TrimSpace(names[1])
		}

		fp := path.Join("file", "music", f.Filename)
		if err := c.SaveUploadedFile(f, fp); err != nil {
			api.ResponseError(c, message.FileUploadFailed)
			return
		}

		musics = append(musics, &domain.Music{
			UID:    c.GetInt64("uid"),
			Name:   name,
			Author: author,
			URL:    fp,
		})
	}

	if err := mh.musicSrv.AddMusics(c, musics); err != nil {
		api.ResponseError(c, message.Failed)
		return
	}

	api.ResponseOK(c, nil)
}

// UpdateMusic 更新音乐
//
//	@receiver mh
//	@param c
func (mh *MusicHandler) UpdateMusic(c *gin.Context) {
	type Req struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	var req Req
	if err := c.ShouldBind(&req); err != nil || req.ID == 0 {
		api.ResponseError(c, message.BadRequest)
		return
	}

	if err := mh.musicSrv.UpdateMusicInfo(c, &domain.Music{
		ID:     req.ID,
		Name:   req.Name,
		Author: req.Author,
	}); err != nil {
		if errors.Is(err, service.ErrMusicFound) {
			api.ResponseError(c, message.MusicNotFound)
			return
		}
		api.ResponseError(c, message.Failed)
		return
	}

	api.ResponseOK(c, nil)
}
