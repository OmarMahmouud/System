package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string   `json:"name" validate:"required" gorm:"type:varchar(50);unique"`
	Description string   `json:"description" validate:"required"  gorm:"type:varchar(500)"`
	Courses     []Course `json:"courses"`
}
