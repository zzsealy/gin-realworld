package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() (*gorm.DB, error) {
    dsn := "drq:drq12345.@tcp(127.0.0.1:3456)/go_learn?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
	DB = db
    return db, nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() {
	  
	DB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	DB.Close()
}





