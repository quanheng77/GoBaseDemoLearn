package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", func(context *gin.Context) {

		name := context.Param("name")
		action := context.Param("action")
		//截取
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})

	r.Run(":8080")
}
