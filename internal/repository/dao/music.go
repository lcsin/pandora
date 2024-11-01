package dao

import (
	"context"

	"gorm.io/gorm"
)

// IMusicDAO music dao interface
type IMusicDAO interface {
	SelectMusicInfoByID(ctx context.Context, ID int64) (*Music, error)

	SelectMusicList(ctx context.Context) ([]*Music, error)

	SelectMyMusicByNameOrAuthor(ctx context.Context, query string) ([]*Music, error)

	InsertMusics(ctx context.Context, musics []Music) error

	UpdateMusicInfo(ctx context.Context, music Music) error

	DeleteMusicByID(ctx context.Context, ID int64) error
}

// Music 数据库音乐表实体映射
type Music struct {
	ID     int64 `gorm:"primaryKey,authIncrement"`
	Name   string
	Author string
	URL    string
}

// TableName 数据库音乐表表名映射
//
//	@receiver m
//	@return string
func (m *Music) TableName() string {
	return "music_tbl"
}

// MusicMySQL music dao
type MusicMySQL struct {
	db *gorm.DB
}

// NewMusicMySQL 音乐DAO构造函数
//
//	@param db
//	@return IMusicDAO
func NewMusicMySQL(db *gorm.DB) IMusicDAO {
	return &MusicMySQL{db: db}
}

// SelectMusicInfoByID 根据ID获取音乐信息
//
//	@receiver m
//	@param ctx
//	@param ID
//	@return *Music
//	@return error
func (m *MusicMySQL) SelectMusicInfoByID(ctx context.Context, ID int64) (*Music, error) {
	var music Music
	if err := m.db.Where("id = ?", ID).First(&music).Error; err != nil {
		return nil, err
	}
	return &music, nil
}

// SelectMusicList 获取音乐列表
//
//	@receiver m
//	@param ctx
//	@param id
//	@return []*Music
//	@return error
func (m *MusicMySQL) SelectMusicList(ctx context.Context) ([]*Music, error) {
	var music []*Music
	if err := m.db.Find(&music).Error; err != nil {
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
func (m *MusicMySQL) SelectMyMusicByNameOrAuthor(ctx context.Context, query string) ([]*Music, error) {
	var music []*Music

	queryScope := func(db *gorm.DB) *gorm.DB {
		// db = db.Where("uid = ?", uid)
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
func (m *MusicMySQL) InsertMusics(ctx context.Context, musics []Music) error {
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
func (m *MusicMySQL) UpdateMusicInfo(ctx context.Context, music Music) error {
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
func (m *MusicMySQL) DeleteMusicByID(ctx context.Context, ID int64) error {
	return m.db.Where("id = ?", ID).Delete(&Music{}).Error
}
