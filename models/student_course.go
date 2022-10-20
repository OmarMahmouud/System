package models

import "gorm.io/gorm"

type StudentCourse struct {
	gorm.Model
	UserID   uint   `json:"user_id" validate:"required" `
	CourseID uint   `json:"course_id" validate:"required"`
	User     User   `json:"user"`
	Course   Course `json:"course"`
}
