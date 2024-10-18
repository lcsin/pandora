package service

import (
	"context"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository"
)

// IMusicService 音乐Service层接口
type IMusicService interface {
	GetMusicListByUID(ctx context.Context, uid int64) ([]*domain.Music, error)
}

// MusicService music service
type MusicService struct {
	repo repository.IMusicRepository
}

// NewMusicService music service构造函数
//
//	@param repo
//	@return *MusicService
func NewMusicService(repo repository.IMusicRepository) IMusicService {
	return &MusicService{repo: repo}
}

// GetMusicListByUID 根据用户id获取音乐列表
//
//	@receiver ms
//	@param ctx
//	@param uid
//	@return []*domain.Music
//	@return error
func (ms *MusicService) GetMusicListByUID(ctx context.Context, uid int64) ([]*domain.Music, error) {
	return ms.repo.GetMusicListByUID(ctx, uid)
}
