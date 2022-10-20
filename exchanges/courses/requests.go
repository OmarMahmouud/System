package courses

import "time"

type StoreRequest struct {
	CourseName  string    `json:"course_name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CategoryID  uint      `json:"category_id"`
}

type IdRequest struct {
	ID uint `json:"id"`
}

type UpdatedRequest struct {
	ID          uint      `json:"id"`
	CourseName  string    `json:"course_name" `
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CategoryID  uint      `json:"category_id"`
}

type GetAllRequest struct {
	CourseName string    `json:"course_name" `
	CategoryID uint      `json:"category_id"`
	UserID     uint      `json:"user_id"`
	StartDate  time.Time `json:"start_date"`
}
