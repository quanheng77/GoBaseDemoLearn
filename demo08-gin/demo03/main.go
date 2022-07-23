package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:9000/user 才会打印出来默认的值
		query := c.DefaultQuery("name", "张三")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", query))
	})
	r.Run(":9000")
}
