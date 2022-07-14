package main

import "fmt"

// Demo 01  hello world 案例

func main() {
	fmt.Println("Hello World,Golang!!!")
	//Hello World,Golang!!!
	var (
		name string
		age  int
	)
	fmt.Scan(&name)
	fmt.Printf("hello %s", name)
	//quanheng
	//hello quanheng
	fmt.Scanf("1:%s 2:%d", &name, &age)
	fmt.Printf("name:%s,age:%d", name, age)
	//1:小王子 2:28
	//name:小王子,age:28
	fmt.Scanln(&name, &age)
	fmt.Printf("name:%s,age:%d", name, age)
	//quanheng 11
	//name:uanheng,age:11
}

/** fmt 实例化I/O 实现了printf 和 scan
General：
	%v :当前值
	%T :当前值属性
Boolean：
	%t: 判断对错
Integer:
	%d	十进制地址值
	%o	八进制地址值
	%O	0o为前缀的地址值
string and slice:
	%s 打印字符串或切面
scan :
	函数扫描标准输入中给定的输入文本，从那里读取内容并将连续的以空格分隔的值存储到连续的参数中
Scanf:
	从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中
Scanln:
	类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置
print:
	Print系列函数会将内容输出到系统的标准输出，区别在于
		Print函数直接输出内容
		Printf函数支持格式化输出字符串
		Println函数会在输出内容的结尾添加一个换行符
	Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
	Sprint系列函数会把传入的数据生成并返回一个字符串
	Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误

*/
