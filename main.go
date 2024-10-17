package main

import (
	"fmt"

	"github.com/lcsin/pandora/ioc"
	"github.com/spf13/viper"
)

func main() {
	ioc.InitConfig()
	r := InitApp()
	r.LoadHTMLGlob("html/*")

	r.Run(fmt.Sprintf(":%v", viper.Get("app.port")))
}
