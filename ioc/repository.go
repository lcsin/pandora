package ioc

import (
	"log"
	"os"
	"time"

	"github.com/lcsin/pandora/internal/repository/dao"
	"gorm.io/driver/sqlite"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitMySQL 初始化mysql
func InitMySQL() *gorm.DB {
	dns := viper.Get("mysql.dns").(string)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&dao.Music{})

	return db
}

// InitSQLite 初始化sqllite
func InitSQLite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("pandora.db"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&dao.Music{})

	return db
}
