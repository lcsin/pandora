package ioc

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// IniDB 初始化mysql
func IniDB() *gorm.DB {
	dns := viper.Get("mysql.dns").(string)
	db, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	return db
}
