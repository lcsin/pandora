package domain

import "github.com/lcsin/pandora/internal/repository/dao"

// User domain.User
type User struct {
	ID          int64
	Email       string
	Username    string
	Password    string
	CreatedTime string
	UpdatedTime string
}

// ToEntity 转换为数据库的User模型
//  @receiver u
//  @return dao.User
func (u *User) ToEntity() dao.User {
	return dao.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedTime,
		UpdatedAt: u.UpdatedTime,
	}
}
