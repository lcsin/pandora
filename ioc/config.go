package ioc

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// InitConfig 初始化配置信息
func InitConfig() {
	fp := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigFile(*fp)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
