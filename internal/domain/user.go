package domain

import (
	"time"

	"github.com/lcsin/pandora/internal/repository/dao"
)

// User domain.User
type User struct {
	ID          int64
	Email       string
	Username    string
	Password    string
	CreatedTime string
	UpdatedTime string
}

// ToDAO 转换为数据库的User模型
//
//	@receiver u
//	@return dao.User
func (u *User) ToDAO() dao.User {
	createdAt, _ := time.Parse(time.DateTime, u.CreatedTime)
	updatedAt, _ := time.Parse(time.DateTime, u.UpdatedTime)

	return dao.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
