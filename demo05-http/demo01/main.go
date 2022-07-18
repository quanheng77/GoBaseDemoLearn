package main

import (
	"fmt"
	"net/http"
)

type User struct {
	username string
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你好"+u.username)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello user!"))
}

func main() {
	http.HandleFunc("/user", UserHandler)
	http.Handle("/username", &User{username: "张三"})
	http.ListenAndServe(":10086", nil)
}

/*
http.HandleFunc 和 http.Handle 都是用于注册路由，可以发现两者的区别在于第二个参数，
前者是一个具有 func(w http.ResponseWriter, r *http.Requests) 签名的函数，而后者是一个结构体，
该结构体实现了 func(w http.ResponseWriter, r *http.Requests) 签名的方法。


Handler 是一个接口
HandlerFunc 实际上是将 handler 函数做了一个类型转换

*/

//https://learnku.com/articles/37867
//https://blog.csdn.net/zrg3699/article/details/122280399
