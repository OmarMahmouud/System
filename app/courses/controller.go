package courses

import (
	"System/System"
	"System/exchanges/courses"
	"System/exchanges/user"
	transformerCourse "System/transformers"
	"github.com/gin-gonic/gin"
)

func storeControl(c *gin.Context) {
	var request courses.StoreRequest
	if valid := storeValidete(c, &request); !valid {
		return
	}
	row := storeService(request)
	res := transformerCourse.TranformCourse(&row, nil)
	System.Created(c, res)
}

func getAllControl(c *gin.Context) {
	var req courses.GetAllRequest
	if valid := jsonValidate(c, &req); !valid {
		return
	}
	AllCourses, studentsCourses := getAllService(&req)
	resS := transformerCourse.TranformCourses(AllCourses, studentsCourses)
	System.Ok(c, resS)
}

func showByIdControl(c *gin.Context) {
	var request courses.IdRequest
	valid, row := idValidate(c, &request)
	if !valid {
		return
	}
	courseStudents := showByIDService(row)
	res := transformerCourse.TranformCourse(row, courseStudents)
	System.Ok(c, res)
}

func deleteControl(c *gin.Context) {
	var request courses.IdRequest
	valid, row := idValidate(c, &request)
	if !valid {
		return
	}
	deleteService(row)
	System.Deleted(c)
}

func updatedControl(c *gin.Context) {
	var request courses.UpdatedRequest
	valid, row := updateValidate(c, &request)
	if !valid {
		return
	}
	updateService(request, row)
	courseStudents := showByIDService(row)
	res := transformerCourse.TranformCourse(row, courseStudents)
	System.Ok(c, res)
}

func StudentJoinCourseControl(c *gin.Context) {
	var req user.Join
	if !validateJoinRequest(c, &req) {
		return
	}
	row := StudentJoinService(req)
	res := transformerCourse.TranformJoin(&row)
	System.Ok(c, res)
}

func StudentDropCourseControl(c *gin.Context) {
	var req user.Join
	if !validateDropRequest(c, &req) {
		return
	}
	DropJoinService(&req)
	System.Deleted(c)
}
