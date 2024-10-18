package dao

import (
	"context"

	"gorm.io/gorm"
)

// IMusicDAO music dao interface
type IMusicDAO interface {
	SelectMusicListByUID(ctx context.Context, uid int64) ([]*Music, error)
}

// Music 数据库音乐表实体映射
type Music struct {
	ID     int64 `gorm:"primaryKey,authIncrement"`
	UID    int64
	Name   string
	Author string
	URL    string
	Time string
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

// SelectMusicListByID 根据用户ID获取音乐列表
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
