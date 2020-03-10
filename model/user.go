package model

import (
	"goedu/utils"

	"github.com/jinzhu/gorm"
)

// User 用户表
type User struct {
	ID        uint    `json:"id"`
	Username  string  `form:"username" json:"username"`
	Email     string  `form:"email" json:"email"`
	Password  string  `form:"password" json:"password"`
	CreatedAt MyTime  `json:"created_at"`
	UpdatedAt MyTime  `json:"updated_at"`
	DeletedAt *MyTime `json:"deleted_at" sql:"index"`
}

// BeforeCreate 创建之前的一些 Hook 操作
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_ = scope.SetColumn("password", hash)
	return nil
}
