package users

import (
	"gin_learn/common"
)

type UserModel struct {
	ID				uint `gorm:"primary_key"`
	Username		string `gorm:"column:username"`
	Email			string `gorm:"column:email;unique_index"`
	PasswordHash 	string `gorm:"column:passwordHash"`
}

func (u *UserModel) TableName() string{
	return "user"
}


func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&UserModel{})
}