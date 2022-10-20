package routerSystem

import (
	"System/app/categories"
	"System/app/courses"
	"System/app/users"
	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	courses.CoursesRoutesByAdmin(r)
	courses.CoursesRoutes(r)
	categories.CategoryRoutes(r)
	users.UserRoutes(r)
}
