package dao

import (
	"context"

	"gorm.io/gorm"
)

// IMusicDAO music dao interface
type IMusicDAO interface {
	SelectMusicInfoByID(ctx context.Context, ID int64) (*Music, error)

	SelectMusicListByUID(ctx context.Context, uid int64) ([]*Music, error)

	SelectMyMusicByNameOrAuthor(ctx context.Context, uid int64, query string) ([]*Music, error)

	InsertMusics(ctx context.Context, musics []Music) error

	UpdateMusicInfo(ctx context.Context, music Music) error

	DeleteMusicByID(ctx context.Context, ID int64) error
}

// Music 数据库音乐表实体映射
type Music struct {
	ID     int64 `gorm:"primaryKey,authIncrement"`
	UID    int64
	Name   string
	Author string
	URL    string
	Time   float64
}

// TableName 数据库音乐表表名映射
//
//	@receiver m
//	@return string
func (m *Music) TableName() string {
	return "music_tbl"
}

// MusicDAO music dao
type MusicDAO struct {
	db *gorm.DB
}

// NewMusicDAO 音乐DAO构造函数
//
//	@param db
//	@return IMusicDAO
func NewMusicDAO(db *gorm.DB) IMusicDAO {
	return &MusicDAO{db: db}
}

// SelectMusicInfoByID 根据ID获取音乐信息
//
//	@receiver m
//	@param ctx
//	@param ID
//	@return *Music
//	@return error
func (m *MusicDAO) SelectMusicInfoByID(ctx context.Context, ID int64) (*Music, error) {
	var music Music
	if err := m.db.Where("id = ?", ID).First(&music).Error; err != nil {
		return nil, err
	}
	return &music, nil
}

// SelectMusicListByUID 根据用户ID获取音乐列表
//
//	@receiver m
//	@param ctx
//	@param id
//	@return []*Music
//	@return error
func (m *MusicDAO) SelectMusicListByUID(ctx context.Context, uid int64) ([]*Music, error) {
	var music []*Music
	if err := m.db.Where("uid = ?", uid).Find(&music).Error; err != nil {
		return nil, err
	}
	return music, nil
}

// SelectMyMusicByNameOrAuthor 根据歌名或作者名获取音乐列表
//
//	@receiver m
//	@param ctx
//	@param name
//	@param author
//	@return []*Music
//	@return error
func (m *MusicDAO) SelectMyMusicByNameOrAuthor(ctx context.Context, uid int64, query string) ([]*Music, error) {
	var music []*Music

	queryScope := func(db *gorm.DB) *gorm.DB {
		db = db.Where("uid = ?", uid)
		if len(query) > 0 {
			db = db.Where("`name` = ? or author = ?", query, query)
		}
		return db
	}

	if err := m.db.Scopes(queryScope).Find(&music).Error; err != nil {
		return nil, err
	}

	return music, nil
}

// InsertMusics 新增音乐
//
//	@receiver m
//	@param ctx
//	@param music
//	@return error
func (m *MusicDAO) InsertMusics(ctx context.Context, musics []Music) error {
	if len(musics) > 0 {
		return m.db.Create(&musics).Error
	}
	return nil
}

// UpdateMusicInfo 更新音乐信息
//
//	@receiver m
//	@param ctx
//	@param music
//	@return error
func (m *MusicDAO) UpdateMusicInfo(ctx context.Context, music Music) error {
	return m.db.Model(&Music{}).Where("id = ?", music.ID).Updates(map[string]interface{}{
		"name":   music.Name,
		"author": music.Author,
	}).Error
}

// DeleteMusicByID 根据ID删除音乐
//
//	@receiver m
//	@param ctx
//	@param ID
//	@return error
func (m *MusicDAO) DeleteMusicByID(ctx context.Context, ID int64) error {
	return m.db.Where("id = ?", ID).Delete(&Music{}).Error
}
