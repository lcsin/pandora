package service

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/lcsin/pandora/internal/domain"
	"github.com/lcsin/pandora/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	// ErrUserNotFound 用户不存在或密码错误
	ErrUserNotFound = errors.New("用户不存在或密码错误")

	// ErrUserExisted 用户已存在
	ErrUserExisted = errors.New("用户已存在")
)

// IUserService UserService Interface
type IUserService interface {
	GetByID(ctx context.Context, uid int64) (*domain.User, error)

	Login(ctx context.Context, email, password string) (*domain.User, error)

	Regiser(ctx context.Context, user domain.User) error
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

// GetByID 根据用户ID获取用户信息
//
//	@receiver us
//	@param ctx
//	@param uid
//	@return *domain.User
//	@return error
func (us *UserService) GetByID(ctx context.Context, uid int64) (*domain.User, error) {
	user, err := us.repo.GetByID(ctx, uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// Login 用户登录
//
//	@receiver us
//	@param ctx
//	@param email
//	@param password
//	@return error
func (us *UserService) Login(ctx context.Context, email string, password string) (*domain.User, error) {
	user, err := us.repo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// Regiser 用户注册
//
//	@receiver us
//	@param ctx
//	@param user
//	@return error
func (us *UserService) Regiser(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.CreatedTime = time.Now().Format(time.DateTime)
	if err = us.repo.Create(ctx, user); err != nil {
		// 判断邮箱是否已经被注册
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrUserExisted
		}
		return err
	}

	return nil
}
