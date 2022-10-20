package categories

import (
	"System/auth"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	r := router.Group("/category").Use(auth.Auth).Use(auth.AuthAdmin)
	r.POST("/store", storeControl)
	r.GET("/all", getAllControl)
	r.GET("/show", ShowByIdControl)
	r.DELETE("/delete", deleteControl)
	r.PUT("/update", updatedControl)
}
