package users

import (
	"System/System"
	"System/app/common"
	"System/app/courses"
	"System/exchanges/user"
	"System/models"
	"github.com/gin-gonic/gin"
)

func loginValidtae(c *gin.Context, req *user.LoginRequest) (bool, models.User) {
	if !jsonValidate(c, &req) {
		return false, models.User{}
	}
	found, row := getUserByEmail(req.Email)
	if !found {
		System.NotFound(c, "We Not found This User")
		return false, models.User{}
	}
	err := VerifyPassword(row.Password, req.Password)
	if !common.CheckError(err) {
		System.BadRequest(c, "enter your the correct password")
		return false, models.User{}
	}
	return true, row
}

func jsonValidate(c *gin.Context, request interface{}) bool {
	err := c.ShouldBindJSON(request)
	if !common.CheckError(err) {
		System.BadRequest(c, err)
		return false
	}
	return true
}

func logoutValidate(c *gin.Context, req *user.LogoutRequest) *models.User {
	if !jsonValidate(c, req) {
		//fmt.Println("aaaaaaaaaaaaaaaaaaa")
		return nil
	}
	row := courses.UserIsFound(c, req.Id)
	if row == nil {
		return nil
	}
	return row
}

func signUpValidate(c *gin.Context, req *user.SignupRequest) bool {
	if check, err := user.ValidStore(c.Request, req); !check {
		System.BadRequest(c, err)
		return false
	}
	return true
}
