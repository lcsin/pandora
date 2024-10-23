package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// IUserDAO 用户DAO层接口
type IUserDAO interface {
	SelectUserByEmail(ctx context.Context, email string) (*User, error)

	Inser(ctx context.Context, user User) error

	UpdateUserByID(ctx context.Context, user User) error
}

// User 数据库用户表实体映射
type User struct {
	ID        int64  `gorm:"primaryKey,autoIncrement"`
	Email     string `gorm:"unique"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 数据库用户表表名映射
func (u *User) TableName() string {
	return "user_tbl"
}

// UserDAO 用户DAO层
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 用户DAO构造函数
func NewUserDAO(db *gorm.DB) IUserDAO {
	return &UserDAO{db: db}
}

// SelectUserByEmail 根据邮箱获取用户信息
//
//	@receiver u
//	@param ctx
//	@param email
//	@return User
//	@return error
func (u *UserDAO) SelectUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Inser 新增用户
//
//	@receiver u
//	@param ctx
//	@param user
//	@return error
func (u *UserDAO) Inser(ctx context.Context, user User) error {
	return u.db.Create(&user).Error
}

// UpdateUserByID 根据用户ID更新用户信息
//
//	@receiver u
//	@param ctx
//	@param user
//	@return error
func (u *UserDAO) UpdateUserByID(ctx context.Context, user User) error {
	return u.db.Updates(&user).Error
}
