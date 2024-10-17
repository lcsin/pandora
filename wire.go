//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/lcsin/pandora/internal/handler"
	"github.com/lcsin/pandora/internal/repository"
	"github.com/lcsin/pandora/internal/repository/dao"
	"github.com/lcsin/pandora/internal/service"
	"github.com/lcsin/pandora/ioc"
)

// InitApp wire 依赖注入
//
//	@return *gin.Engine
func InitApp() *gin.Engine {
	wire.Build(
		// 数据库、web服务
		ioc.InitDB, ioc.InitWebServer, handler.NewWebHandler,
		// 用户服务
		dao.NewUserDAO, repository.NewUserRepository, service.NewUserService, handler.NewUserHandler,
	)

	return gin.Default()
}
