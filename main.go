package main

import (
	"fmt"
	"gin_learn/common"
	"gin_learn/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	fmt.Println("migrate db", db)
}

func main() {
	db, err := common.Init()
	if err != nil {
		fmt.Println("连接数据库发生错误：", err)
	} else {
		fmt.Println("数据库连接成功!")
	}
	Migrate(db)
	common.CloseDB()

	r := gin.Default()
	// v1 = r.Group("/api")

	r.Run()
}