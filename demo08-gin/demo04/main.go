package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		typeName := c.DefaultQuery("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,typeName:%s", username, password, typeName))
	})
	r.Run(":8080")
}
