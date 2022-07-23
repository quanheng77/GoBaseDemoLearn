package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	//3.监听端口
	go r.Run("0.0.0.0:8080")
	go r.Run("127.0.0.1:8081")
	select {}
}

//Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
//对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错
//借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

//go get -u github.com/gin-gonic/gin 安装命令
