package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository"
	"gorm.io/gorm"
)

var (
	// ErrMusicFound 音乐不存在
	ErrMusicFound = errors.New("音乐不存在")
)

// IMusicService 音乐Service层接口
type IMusicService interface {
	GetMusicInfoByID(ctx context.Context, ID int64) (*domain.Music, error)

	GetMusicList(ctx context.Context) ([]*domain.Music, error)

	GetMyMusicListByNameOrAuthor(ctx context.Context, query string) ([]*domain.Music, error)

	AddMusics(ctx context.Context, musics []*domain.Music) error

	UpdateMusicInfo(ctx context.Context, music *domain.Music) error

	DeleteMusicByID(ctx context.Context, ID int64) error
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

// GetMusicInfoByID 根据ID获取音乐信息
//
//	@receiver ms
//	@param ctx
//	@param ID
//	@return *domain.Music
//	@return error
func (ms *MusicService) GetMusicInfoByID(ctx context.Context, ID int64) (*domain.Music, error) {
	return ms.repo.GetMusicInfoByID(ctx, ID)
}

// GetMusicList 获取音乐列表
//
//	@receiver ms
//	@param ctx
//	@param uid
//	@return []*domain.Music
//	@return error
func (ms *MusicService) GetMusicList(ctx context.Context) ([]*domain.Music, error) {
	return ms.repo.GetMusicList(ctx)
}

// GetMyMusicListByNameOrAuthor 根据歌名或作者获取音乐列表
//
//	@receiver ms
//	@param ctx
//	@param name
//	@param author
//	@return []*domain.Music
//	@return error
func (ms *MusicService) GetMyMusicListByNameOrAuthor(ctx context.Context, query string) ([]*domain.Music, error) {
	return ms.repo.GetMyMusicListByNameOrAuthor(ctx, query)
}

// AddMusics 新增音乐
//
//	@receiver ms
//	@param ctx
//	@param music
//	@return error
func (ms *MusicService) AddMusics(ctx context.Context, musics []*domain.Music) error {
	return ms.repo.AddMusics(ctx, musics)
}

// UpdateMusicInfo 更新音乐信息
//
//	@receiver ms
//	@param ctx
//	@param music
//	@return error
func (ms *MusicService) UpdateMusicInfo(ctx context.Context, music *domain.Music) error {
	if music == nil {
		return fmt.Errorf("music is nil")
	}

	// 判断该音乐是否存在
	_, err := ms.GetMusicInfoByID(ctx, music.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMusicFound
		}
		return err
	}

	// 更新音乐信息
	if err := ms.repo.UpdateMusicInfo(ctx, music); err != nil {
		return err
	}

	return nil
}

// DeleteMusicByID 根据ID删除音乐
//
//	@receiver ms
//	@param ctx
//	@param ID
//	@return error
func (ms *MusicService) DeleteMusicByID(ctx context.Context, ID int64) error {
	// 判断该音乐是否存在
	music, err := ms.GetMusicInfoByID(ctx, ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrMusicFound
		}
		return err
	}

	if err := ms.repo.DeleteMusicByID(ctx, ID); err != nil {
		return err
	}
	// 删除本地文件
	_ = os.Remove(music.URL)

	return nil
}
