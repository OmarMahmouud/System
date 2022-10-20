package main

import (
	"System/System"
	"System/customRule"
	"System/routerSystem"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	customRule.Init()
	System.ConnectToDataBase()
	System.Migrate()
	routerSystem.Routing(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
