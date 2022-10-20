package System

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, code int, body interface{}, message interface{}) {
	c.JSON(code, gin.H{
		"data":    body,
		"message": message,
	})
}
func Ok(c *gin.Context, body interface{}) {
	Response(c, http.StatusOK, body, "ok")
}

func Created(c *gin.Context, body interface{}) {
	Response(c, http.StatusCreated, body, "Created")
}

func BadRequest(c *gin.Context, message interface{}) {
	Response(c, http.StatusBadRequest, nil, message)
}
func NotFound(c *gin.Context, message string) {
	Response(c, http.StatusNotFound, nil, message)
}

func Deleted(c *gin.Context) {
	Response(c, http.StatusOK, "Deleted", "ok")
}

func NotUnauthorized(c *gin.Context) {
	Response(c, http.StatusUnauthorized, nil, "not_unauthorized")
}
