package transformers

import (
	"System/exchanges/category"
	"System/exchanges/courses"
	"System/exchanges/user"
	"System/models"
)

func tranformCatOfCourse(cat models.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:          cat.ID,
		Name:        cat.Name,
		Description: cat.Description,
	}
}
func tranformStudentOfCourse(Student models.StudentCourse) (student user.Student) {
	return user.Student{
		ID:   Student.User.ID,
		Name: Student.User.Name,
	}
}
func tranformStudentsOfCourse(studentsCourse []models.StudentCourse) (students []user.Student) {
	for _, student := range studentsCourse {
		students = append(students, tranformStudentOfCourse(student))
	}
	return
}
func TranformCourse(course *models.Course, students []models.StudentCourse) courses.CourseResponse {
	return courses.CourseResponse{
		CourseName:  course.CourseName,
		Description: course.Description,
		StartDate:   course.StartDate,
		EndDate:     course.EndDate,
		ID:          course.ID,
		Category:    tranformCatOfCourse(course.Category),
		Students:    tranformStudentsOfCourse(students),
	}
}

func TranformCourses(courses []models.Course, students []models.StudentCourse) (responses []courses.CourseResponse) {
	studentMap := buildStudentMap(students)
	for _, course := range courses {
		responses = append(responses, TranformCourse(&course, studentMap[course.ID]))
	}
	return
}

func buildStudentMap(students []models.StudentCourse) map[uint][]models.StudentCourse {
	studentMap := make(map[uint][]models.StudentCourse)
	for _, student := range students {
		studentMap[student.CourseID] = append(studentMap[student.CourseID], student)
	}
	return studentMap
}
