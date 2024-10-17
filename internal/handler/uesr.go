package handler

import "github.com/lcsin/pandora/internal/service"

// UserHandler user handler
type UserHandler struct {
	userSrv service.IUserService
}

// NewUserHandler user handler 构造器
//  @param userSrv
//  @return *UserHandler
func NewUserHandler(userSrv service.IUserService) *UserHandler {
	return &UserHandler{userSrv: userSrv}
}
