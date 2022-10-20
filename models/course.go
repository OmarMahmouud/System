package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	gorm.Model
	CourseName  string    `json:"course_name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CategoryID  uint      `json:"category_id" validate:"required"`
	Category    Category  `json:"category"`
}
