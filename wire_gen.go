// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/internal/handler"
	"github.com/lcsin/pandora/internal/repository"
	"github.com/lcsin/pandora/internal/repository/dao"
	"github.com/lcsin/pandora/internal/service"
	"github.com/lcsin/pandora/ioc"
)

// Injectors from wire.go:

// InitApp wire 依赖注入
//
//	@return *gin.Engine
func InitApp() *gin.Engine {
	db := ioc.InitDB()
	iUserDAO := dao.NewUserDAO(db)
	iUserRepository := repository.NewUserRepository(iUserDAO)
	iUserService := service.NewUserService(iUserRepository)
	userHandler := handler.NewUserHandler(iUserService)
	engine := ioc.InitWebServer(userHandler)
	return engine
}
