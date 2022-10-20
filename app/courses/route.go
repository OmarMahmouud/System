package courses

import (
	"System/auth"
	"github.com/gin-gonic/gin"
)

func CoursesRoutesByAdmin(router *gin.Engine) {
	r := router.Group("/course").Use(auth.Auth).Use(auth.AuthAdmin)
	r.POST("/store", storeControl)
	r.GET("/all", getAllControl)
	r.GET("/show", showByIdControl)
	r.DELETE("/delete", deleteControl)
	r.PUT("/update", updatedControl)
}

func CoursesRoutes(router *gin.Engine) {
	r := router.Group("/course").Use(auth.Auth)
	r.POST("/join", StudentJoinCourseControl)
	r.DELETE("/drop", StudentDropCourseControl)
}
