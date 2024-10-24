package repository

import (
	"context"
	"time"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository/dao"
)

// IUserRepository 用户Repository层接口
type IUserRepository interface {
	GetByID(ctx context.Context, uid int64) (*domain.User, error)

	GetByEmail(ctx context.Context, email string) (*domain.User, error)

	Create(ctx context.Context, user domain.User) error

	UpdateByID(ctx context.Context, user domain.User) error
}

// UserRepository 用户Reposiroty
type UserRepository struct {
	dao dao.IUserDAO
}

// NewUserRepository 用户Repository构造函数
//
//	@param dao
//	@return IUserRepository
func NewUserRepository(dao dao.IUserDAO) IUserRepository {
	return &UserRepository{dao: dao}
}

// GetByID 根据用户ID获取用户信息
//
//	@receiver ur
//	@param ctx
//	@param uid
//	@return *domain.User
//	@return error
func (ur *UserRepository) GetByID(ctx context.Context, uid int64) (*domain.User, error) {
	user, err := ur.dao.SelectUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		CreatedTime: user.CreatedAt.Format(time.DateTime),
		UpdatedTime: user.CreatedAt.Format(time.DateTime),
	}, nil
}

// Create 创建用户
//
//	@receiver ur
//	@param ctx
//	@param user
//	@return error
func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	return ur.dao.Inser(ctx, user.ToDAO())
}

// GetByEmail 根据email获取用户
//
//	@receiver ur
//	@param ctx
//	@param email
//	@return domain.User
//	@return error
func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := ur.dao.SelectUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		CreatedTime: user.CreatedAt.Format(time.DateTime),
		UpdatedTime: user.UpdatedAt.Format(time.DateTime),
	}, nil
}

// UpdateByID 根据用户ID更新用户信息
//
//	@receiver ur
//	@param ctx
//	@param user
//	@return error
func (ur *UserRepository) UpdateByID(ctx context.Context, user domain.User) error {
	return ur.dao.UpdateUserByID(ctx, user.ToDAO())
}
