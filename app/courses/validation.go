package courses

import (
	"System/System"
	"System/app/categories"
	"System/app/common"
	"System/exchanges/courses"
	"System/exchanges/user"
	"System/models"
	"github.com/gin-gonic/gin"
)

func jsonValidate(g *gin.Context, request interface{}) bool {
	err := g.ShouldBindJSON(request)
	if !common.CheckError(err) {
		System.BadRequest(g, err)
		return false
	}
	return true
}

func isFoundValidate(c *gin.Context, id uint) (bool, *models.Course) {
	row := GetCourseByID(id)
	if row.ID == 0 {
		System.NotFound(c, "We Not found This Course")
		return false, nil
	}
	return true, &row
}

func storeValidete(c *gin.Context, request *courses.StoreRequest) bool {
	if check, err := courses.ValidStore(c.Request, request); !check {
		System.BadRequest(c, err)
		return false
	}
	if categories.GetCategoryByID(request.CategoryID).ID == 0 {
		System.BadRequest(c, "This Category is not found")
		return false
	}
	return true
}

func idValidate(c *gin.Context, request *courses.IdRequest) (bool, *models.Course) {
	if !jsonValidate(c, request) {
		return false, nil
	}
	if found, row := isFoundValidate(c, request.ID); found {
		return true, row
	}
	return false, nil
}

func updateValidate(c *gin.Context, request *courses.UpdatedRequest) (bool, *models.Course) {
	if check, err := courses.ValidUbdate(c.Request, request); !check {
		System.BadRequest(c, err)
		return false, nil
	}
	if found, row := isFoundValidate(c, request.ID); found {
		return true, row
	}
	return false, nil
}

func courseIsFound(c *gin.Context, id uint) uint {
	row := GetCourseByID(id)
	if row.ID == 0 {
		System.NotFound(c, "We Not found This Course")
		return 0
	}
	return id
}
func UserIsFound(c *gin.Context, id uint) *models.User {
	row := getUserByID(id)
	if row.ID == 0 {
		System.NotFound(c, "We Not found This User")
		return nil
	}
	return &row
}
func userIsStudent(c *gin.Context, id uint) bool {
	row := getUserByID(id)
	if row.Role == "Student" {
		return true
	}
	System.BadRequest(c, "not student")
	return false
}

func validateJoinRequest(c *gin.Context, request *user.Join) bool {
	if !jsonValidate(c, request) {
		return false
	}
	if idCourse := courseIsFound(c, request.CourseID); idCourse == 0 {
		return false
	}
	if UserIsFound(c, request.UserID) == nil {
		return false
	}
	if !userIsStudent(c, request.UserID) {
		return false
	}
	if id := JoinValidate(request); id != 0 {
		System.BadRequest(c, "you are Already register")
		return false
	}
	return true
}

func validateDropRequest(c *gin.Context, request *user.Join) bool {
	if !jsonValidate(c, request) {
		return false
	}
	if id := JoinValidate(request); id == 0 {
		System.BadRequest(c, "not found this register")
		return false
	}
	if !userIsStudent(c, request.UserID) {
		return false
	}
	return true
}
