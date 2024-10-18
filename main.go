package main

import (
	"fmt"

	"github.com/lcsin/pandora/ioc"
	"github.com/spf13/viper"
)

func main() {
	ioc.InitConfig()
	r := InitApp()
	// web.RegisterRoutes(r) // 注册前端页面路由

	r.Run(fmt.Sprintf(":%v", viper.Get("app.port")))
}
