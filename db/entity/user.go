package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"username,index,unique"`
	PasswordHash string `json:"-" gorm:"password"`
}

type UserLoginForm struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserRegisterForm struct {
	Username  string `json:"username" form:"username" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	CPassword string `json:"c_password" form:"c_password" validate:"required|eq_field:password"`
}
