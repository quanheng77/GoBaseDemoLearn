package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()
	var login Login
	r.GET("/login", func(c *gin.Context) {
		login.Code = "200"
		login.Msg = "查询成功"

		c.JSON(http.StatusOK, gin.H{
			"msg": "查询成功",
		})
		c.JSON(http.StatusOK, login)
		c.XML(http.StatusOK, login)
		c.YAML(http.StatusOK, login)

	})

	// 5.protobuf格式,谷歌开发的高效存储读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(200, data)
	})
	r.Run()
}

type Login struct {
	Code string `json:"code" xml:"code" yaml:"code"  binding:"code"`
	Msg  string `json:"msg" xml:"msg" yaml:"msg" binding:"msg"`
}
