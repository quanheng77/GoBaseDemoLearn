package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换
	//LoadHTMLGlob()方法可以加载模板文件
	r.LoadHTMLGlob("demo08-gin/demo10/static/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
		//重定向
		//c.Redirect(http.StatusMovedPermanently, "http://www.quanheng77.top")
	})
	r.Run()
}

// https://laravelacademy.org/post/21880 自定义模板
