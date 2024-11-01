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
		ioc.InitSQLLite, ioc.InitWebServer, ioc.InitMiddlewares,
		// 用户服务
		// dao.NewUserDAO, repository.NewUserRepository, service.NewUserService, handler.NewUserHandler,
		// 音乐服务
		dao.NewMusicMySQL, repository.NewMusicRepository, service.NewMusicService, handler.NewMusicHandler,
	)

	return gin.Default()
}
