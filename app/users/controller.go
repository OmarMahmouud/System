package users

import (
	"System/System"
	"System/exchanges/user"
	"System/transformers"
	"github.com/gin-gonic/gin"
)

func loginController(c *gin.Context) {
	var req user.LoginRequest
	found, row := loginValidtae(c, &req)
	if !found {
		return
	}
	loginService(&row)
	res := transformers.TransformUserLogin(&row)
	System.Ok(c, res)
}

func signupControl(c *gin.Context) {
	var req user.SignupRequest
	if !signUpValidate(c, &req) {
		return
	}
	row := signupService(&req)
	res := transformers.TransformUserSignUp(row)
	System.Created(c, res)
}

func logoutControl(c *gin.Context) {
	var req user.LogoutRequest
	row := logoutValidate(c, &req)
	if row == nil {
		return
	}
	logoutService(row)
	System.Ok(c, "you Now Need to signin")
}
