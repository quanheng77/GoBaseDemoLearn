package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", login)
		userGroup.POST("/adduser")
	}
	adminGroup := r.Group("/admin")
	{
		adminGroup.GET("/queryList", login)
		adminGroup.POST("/editList")
	}
	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
