package courses

import (
	"System/System"
	"System/exchanges/courses"
	"System/exchanges/user"
	"System/models"
	"fmt"
)

func store(request courses.StoreRequest) models.Course {
	row := setHelper(request)
	System.DB.Create(&row)
	System.DB.Preload("Category").Where("id = ?", request.CategoryID).Find(&row)
	return row
}

func update(request courses.UpdatedRequest, row *models.Course) {
	updateHelper(request, row)
	System.DB.Save(&row)
	System.DB.Preload("Category").Where("id = ?", row.ID).Find(&row)
}

func GetCourseByID(id uint) models.Course {
	var row models.Course
	System.DB.Preload("Category").Where("id = ?", id).Find(&row)
	return row
}

func all(request *courses.GetAllRequest) []models.Course {
	var rows []models.Course
	System.DB.Preload("Category").Find(&rows)
	if request.CourseName != "" {
		request.CourseName = "%" + request.CourseName + "%"
		fmt.Println(request.CourseName)
		System.DB.Preload("Category").Where("course_name LIKE ?", request.CourseName).Find(&rows)
	}
	if request.CategoryID != 0 {
		System.DB.Preload("Category").Where("category_id = ?", request.CategoryID).Find(&rows)
	}
	if request.UserID != 0 {
		var studentCourses []int
		System.DB.Model(models.StudentCourse{}).Where("user_id = ?", request.UserID).Pluck("course_id", &studentCourses)
		System.DB.Preload("Category").Where("id IN ?", studentCourses).Find(&rows)
	}
	if !request.StartDate.IsZero() {
		System.DB.Preload("Category").Where("start_date = ?", request.StartDate).Find(&rows)
	}
	return rows
}

func deleteCourse(id uint) {
	System.DB.Where("id = ?", id).Unscoped().Delete(&models.Course{})
}

func getStudentsOfCourse(id uint) []models.StudentCourse {
	var rows []models.StudentCourse
	System.DB.Where("course_id = ?", id).Preload("User").Find(&rows)
	return rows
}

func getStudentsOfCourses(ids []uint) []models.StudentCourse {
	var rows []models.StudentCourse
	System.DB.Where("course_id IN ?", ids).Preload("User").Find(&rows)
	return rows
}

func getUserByID(id uint) models.User {
	var row models.User
	System.DB.Where("id = ?", id).First(&row)
	return row
}

func NewJoin(req user.Join) models.StudentCourse {
	row := setJoinHelper(req)
	System.DB.Create(&row)
	return row
}

func JoinValidate(req *user.Join) uint {
	var row models.StudentCourse
	System.DB.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).Unscoped().First(&row)
	return row.ID
}

func DropJoin(req *user.Join) {
	System.DB.Where("user_id = ? AND course_id = ?", req.UserID, req.CourseID).Unscoped().Delete(&models.StudentCourse{})
}
