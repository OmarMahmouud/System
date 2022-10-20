package courses

import (
	"System/app/categories"
	"System/exchanges/courses"
	"System/exchanges/user"
	"System/models"
)

func setHelper(req courses.StoreRequest) models.Course {
	return models.Course{
		CourseName:  req.CourseName,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CategoryID:  req.CategoryID,
	}
}
func updateHelper(req courses.UpdatedRequest, course *models.Course) {
	if req.CourseName != "" {
		course.CourseName = req.CourseName
	}
	if req.Description != "" {
		course.Description = req.Description
	}
	if !req.StartDate.IsZero() {
		course.StartDate = req.StartDate
	}
	if !req.EndDate.IsZero() {
		course.EndDate = req.EndDate
	}
	if req.CategoryID != 0 {
		course.Category = categories.GetCategoryByID(req.CategoryID)
	}
}

func setJoinHelper(req user.Join) models.StudentCourse {
	return models.StudentCourse{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
}
