package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"  gorm:"type:varchar(50)"`
	Role     string `json:"role" gorm:"type:varchar(10);default:'Student'"`
	Email    string `json:"email" validate:"email,required"  gorm:"unique"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
	ISAuth   bool
}
