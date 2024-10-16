package service

import (
	"context"

	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository"
)

// IUserService UserService Interface
type IUserService interface {
	Create(ctx context.Context, user domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateByID(ctx context.Context, user domain.User) error
}

// UserService UserSercie
type UserService struct {
	repo repository.IUserRepository
}

// NewUserService UserService构造器
//
//	@param repo
func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{repo: repo}
}

// Create 创建用户
//
//	@receiver us
//	@param ctx
//	@param user
//	@return error
func (us *UserService) Create(ctx context.Context, user domain.User) error {
	panic("not implemented") // TODO: Implement
}

// GetByEmail 根据用户ID获取用户信息
//
//	@receiver us
//	@param ctx
//	@param email
//	@return *domain.User
//	@return error
func (us *UserService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("not implemented") // TODO: Implement
}

// UpdateByID 根据用户ID更新用户信息
//
//	@receiver us
//	@param ctx
//	@param user
//	@return error
func (us *UserService) UpdateByID(ctx context.Context, user domain.User) error {
	panic("not implemented") // TODO: Implement
}
