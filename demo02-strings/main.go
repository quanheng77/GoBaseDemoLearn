package main

import "strings"

func main() {
	split := strings.Split("2022-01", "-")
	for i := range split {
		println(split[i])
	}
	//2022
	//01
	splitN := strings.SplitN("2022-01-22", "-", 2)
	for i := range splitN {
		println(splitN[i])
	}
	//2022
	//01-22

	fields := strings.Fields("hello world golang")
	for i := range fields {
		println(fields[i])
	}
	//hello
	//world
	//golang

	fieldsFunc := strings.FieldsFunc("helle world", func(r rune) bool {
		//rune是int32的别名，在所有方面都与int32等同。它是 习惯上用于区分字符值和整数值。
		return true
	})
	for i := range fieldsFunc {
		println(fieldsFunc[i])
	}
	//helle world

	after := strings.SplitAfter("hello,java,go,php ", ",")
	for i := range after {
		println(after[i])
	}
	//hello,
	//java,
	//go,
	//php
	n := strings.SplitAfterN("docker,java,go,php ", ",", 2)
	for i := range n {
		println(n[i])
	}
	//docker,
	//java,go,php
}

/* strings go语言的标准库主要用于字符串查找、替换、比较等

切割：
	Split: 切割字符串返回一个切片
	SplitN: 定义一个长度,切割字符串返回一个切片
	Fields: 用空白做分割，返回切面
	FieldsFunc: 满足r == true 的字符串分割，返回一个切片
	SplitAfter: 切割的同时也返回sep,返回一个切片
	SplitAfterN: 定义一个切片长的，切割的同时也返回sep,返回一个切片
*/
//https://zhuanlan.zhihu.com/p/375953005 参考链接
