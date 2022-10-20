package courses

import (
	"System/exchanges/courses"
	"System/exchanges/user"
	"System/models"
)

func storeService(request courses.StoreRequest) models.Course {
	return store(request)
}

func updateService(req courses.UpdatedRequest, row *models.Course) {
	update(req, row)
}

func getAllService(req *courses.GetAllRequest) ([]models.Course, []models.StudentCourse) {
	allCourses := all(req)
	var ids []uint
	for _, course := range allCourses {
		ids = append(ids, course.ID)
	}
	studentsCourses := getStudentsOfCourses(ids)
	return allCourses, studentsCourses
}

func deleteService(row *models.Course) {
	deleteCourse(row.ID)
}

func showByIDService(row *models.Course) []models.StudentCourse {
	return getStudentsOfCourse(row.ID)
}

func StudentJoinService(req user.Join) models.StudentCourse {
	return NewJoin(req)
}
func DropJoinService(req *user.Join) {
	DropJoin(req)
}
