package repository

import (
	"context"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository/dao"
)

// IMusicRepository 音乐Repository层接口
type IMusicRepository interface {
	GetMusicInfoByID(ctx context.Context, ID int64) (*domain.Music, error)

	GetMusicListByUID(ctx context.Context, uid int64) ([]*domain.Music, error)

	GetMyMusicListByNameOrAuthor(ctx context.Context, uid int64, query string) ([]*domain.Music, error)

	AddMusics(ctx context.Context, musics []*domain.Music) error

	UpdateMusicInfo(ctx context.Context, music *domain.Music) error
}

// MusicRpository music repository
type MusicRpository struct {
	dao dao.IMusicDAO
}

// NewMusicRepository 音乐Repository构造函数
//
//	@param dao
//	@return IMusicRepository
func NewMusicRepository(dao dao.IMusicDAO) IMusicRepository {
	return &MusicRpository{dao: dao}
}

// GetMusicInfoByID 根据ID获取音乐信息
//
//	@receiver mr
//	@param ctx
//	@param ID
func (mr *MusicRpository) GetMusicInfoByID(ctx context.Context, ID int64) (*domain.Music, error) {
	music, err := mr.dao.SelectMusicInfoByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &domain.Music{
		ID:     music.ID,
		UID:    music.UID,
		Name:   music.Name,
		Author: music.Author,
		Time:   music.Time,
		URL:    music.URL,
	}, nil
}

// GetMusicListByUID 根据用户id获取音乐列表
//
//	@receiver mr
//	@param ctx
//	@param uid
//	@return []*domain.Music
//	@return error
func (mr *MusicRpository) GetMusicListByUID(ctx context.Context, uid int64) ([]*domain.Music, error) {
	musicList, err := mr.dao.SelectMusicListByUID(ctx, uid)
	if err != nil {
		return nil, err
	}

	music := make([]*domain.Music, 0, len(musicList))
	for _, v := range musicList {
		music = append(music, &domain.Music{
			ID:     v.ID,
			Name:   v.Name,
			Author: v.Author,
			URL:    v.URL,
			Time:   v.Time,
		})
	}

	return music, nil
}

// GetMyMusicListByNameOrAuthor 根据歌名或作者获取音乐列表
//
//	@receiver mr
//	@param ctx
//	@param name
//	@param author
//	@return []*domain.Music
//	@return error
func (mr *MusicRpository) GetMyMusicListByNameOrAuthor(ctx context.Context, uid int64, query string) ([]*domain.Music, error) {
	musicList, err := mr.dao.SelectMyMusicByNameOrAuthor(ctx, uid, query)
	if err != nil {
		return nil, err
	}

	music := make([]*domain.Music, 0, len(musicList))
	for _, v := range musicList {
		music = append(music, &domain.Music{
			ID:     v.ID,
			Name:   v.Name,
			Author: v.Author,
			URL:    v.URL,
			Time:   v.Time,
		})
	}

	return music, nil
}

// AddMusics 新增音乐
//
//	@receiver mr
//	@param ctx
//	@param music
//	@return error
func (mr *MusicRpository) AddMusics(ctx context.Context, musics []*domain.Music) error {
	list := make([]dao.Music, 0, len(musics))
	for _, v := range musics {
		list = append(list, dao.Music{
			UID:    v.UID,
			Name:   v.Name,
			Author: v.Author,
			URL:    v.URL,
			Time:   v.Time,
		})
	}

	return mr.dao.InsertMusics(ctx, list)
}

// UpdateMusicInfo 更新音乐信息
//
//	@receiver mr
//	@param ctx
//	@param music
//	@return error
func (mr *MusicRpository) UpdateMusicInfo(ctx context.Context, music *domain.Music) error {
	return mr.dao.UpdateMusicInfo(ctx, dao.Music{
		ID:     music.ID,
		Name:   music.Name,
		Author: music.Author,
	})
}
