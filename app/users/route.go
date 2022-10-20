package users

import (
	"System/auth"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	r := router.Group("/user")
	r.POST("/signup", signupControl)
	r.POST("/login", loginController)
	r.POST("/logout", logoutControl).Use(auth.Auth)
}
