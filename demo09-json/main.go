package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		// 将request的body中的数据
		// 根据请求头中content-type自动推断
		// Bind()默认解析并绑定form格式
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.Username != "admin" || json.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `form:"username" json:"username" uri:"username" xml:"username" `
	Password string `form:"password" json:"password" uri:"password" xml:"password"`
}
