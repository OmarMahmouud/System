package courses

import (
	"System/exchanges/category"
	"System/exchanges/user"
	"time"
)

type CourseResponse struct {
	ID          uint                      `json:"id"`
	CourseName  string                    `json:"course_name" `
	Description string                    `json:"description"`
	StartDate   time.Time                 `json:"start_date"`
	EndDate     time.Time                 `json:"end_date"`
	Category    category.CategoryResponse `json:"category"`
	Students    []user.Student            `json:"students"`
}
