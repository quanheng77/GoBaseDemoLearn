package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "zhang san")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, fmt.Sprintf("404 not found %s", time.Now().Format("2006-01-02 15:04:05")))
	})
	r.Run()
}
