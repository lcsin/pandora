package repository

import (
	"context"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository/dao"
)

// IMusicRepository 音乐Repository层接口
type IMusicRepository interface {
	GetMusicListByUID(ctx context.Context, uid int64) ([]*domain.Music, error)
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
