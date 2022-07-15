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

	count := strings.Count("shi wo de sha gua", "s")
	println(count)
	//2

	contains := strings.Contains("java,go,docker,kubernetes", "go")
	println(contains)
	//true

	containsAny := strings.ContainsAny("java,go,docker,kubernetes", "kb")
	println(containsAny)
	//true

	index := strings.Index("java,go,docker,kubernetes", "go")
	println(index)
	//5
	lastIndex := strings.LastIndex("java,go,docker", "o")
	println(lastIndex)
	//9

	title := strings.Title("java,go,docker")
	println(title)
	//Java,Go,Docker

	lower := strings.ToLower(title)
	println(lower)
	//java,go,docker

	upper := strings.ToUpper(lower)
	println(upper)
	//JAVA,GO,DOCKER

	space := strings.TrimSpace("  hello gopher  ")
	println(space)
	//hello gopher

	trim := strings.Trim("c,java,go,php,python,pb,c", "c")
	println(trim)
	//,java,go,php,python,pb,

	arr := []string{"go", "java"}
	join := strings.Join(arr, "c")
	println(join)
	//gocjava

	repeat := strings.Repeat("Go", 10)
	println(repeat)
	//GoGoGoGoGoGoGoGoGoGo

	replace := strings.Replace("go,java,c,c,c", "c", "docker", 2)
	println(replace)
	//go,java,docker,docker,c
	all := strings.ReplaceAll("go,java,c,c,c", "c", "docker")
	println(all)
	//go,java,docker,docker,docker

	compare := strings.Compare("hello111", "hello")
	println(compare)
	//1

	fold := strings.EqualFold("gogogo", "go")
	println(fold)
	//false

	builder := strings.Builder{}
	builder.Write([]byte("hello world1"))
	builder.Write([]byte("hello world2"))
	s := builder.String()
	println(s)
	//hello world1hello world2

	reader := strings.Reader{}
	read, _ := reader.Read([]byte("111"))
	println(read)
	//0
}

/* strings go语言的标准库主要用于字符串查找、替换、比较等

切割：
	Split: 切割字符串返回一个切片
	SplitN: 定义一个长度,切割字符串返回一个切片
	Fields: 用空白做分割，返回切面
	FieldsFunc: 满足r == true 的字符串分割，返回一个切片
	SplitAfter: 切割的同时也返回sep,返回一个切片
	SplitAfterN: 定义一个切片长的，切割的同时也返回sep,返回一个切片

查找：
	Count: 字符串包含了多少个所需要查询的字段的个数
	Contains: 是否包含此字符串
	ContainsAny： 是否包含字符串任意一个字符
	Index: 返回字符串首次出现的位置
	LastIndex: 字符串最后一次出现的位置

大小写：
	Title: 每个单词首字母都大写
	ToLower: 都转换为小写
	ToUpper: 都转转化大写

删除：
	TrimSpace：去除两边空白
	Trim: 删除首尾的字符串

拼接去重：
	Join: 拼接数组的字符串
	Repeat: 重复n次

替换：
	Replace: 替换字符串的字符
	ReplaceAll: 全部替换

比较：
	compare: 按照字典比较字符串大小 a>b 返回大于0值 a==b 返回0 a<b 返回小于0值
	EqualFold: 不区分大小写比较两个值的UTF8字符串是否相等

builder :
	写入的四种方式：
	func (b *Builder) Write(p []byte) (int, error)
	func (b *Builder) WriteByte(c byte) error
	func (b *Builder) WriteRune(r rune) (int, error)
	func (b *Builder) WriteString(s string) (int, error)

	我们通过调用 string.Builder 的写入方法来写入内容，然后通过调用 String() 方法来获取拼接的字符串。
	通过使用一个内部的 slice 来存储数据片段。当开发者调用写入方法的时候，数据实际上是被追加（append）到了其内部的 slice 上。

*/

//https://zhuanlan.zhihu.com/p/375953005 strings参考链接
//https://zhuanlan.zhihu.com/p/147042526 builder参考链接
