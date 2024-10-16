package main

import (
	"fmt"
	"time"

	"github.com/lcsin/pandora/internal/repository/dao"
	"github.com/lcsin/pandora/ioc"
)

func main() {
	ioc.InitConfig()
	db := ioc.IniDB()

	var root dao.User
	if err := db.Where("id = 1").First(&root).Error; err != nil {
		panic(err)
	}
	fmt.Println("root=>",root)

	user := dao.User{
		Email: "10086@qq.com",
		Username: "test2",
		Password: "test2",
		CreatedAt: time.Now().Format(time.DateTime),
	}
	if err := db.Create(&user).Error; err != nil {
		panic(err)
	}
}
