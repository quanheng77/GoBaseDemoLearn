package main

import (
	"os"
)

func main() {

	// 获取主机名
	hostname, _ := os.Hostname()
	println(hostname)

	//Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝
	environ := os.Environ()
	for i := range environ {
		println(environ[i])
	}

	//Getenv 函数会检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串
	getenv := os.Getenv("ProgramFiles")
	println(getenv)
	//C:\Program Files

	//Setenv 函数可以设置名为 key 的环境变量，如果出错会返回该错误
	/*err := os.Setenv("GO", "")
	if err != nil{
		log.Fatalf("err:%v\n",err)
	}*/

	//Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
	//os.Exit()

	//Getuid 函数可以返回调用者的用户 ID
	getuid := os.Getuid()
	println(getuid)

	///Getgid 函数可以返回调用者的组 ID。
	getgid := os.Getgid()
	println(getgid)

	///Getpid 函数可以返回调用者所在进程的进程 ID。
	getpid := os.Getpid()
	println(getpid)

	//Getwd 函数可以返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个
	getwd, _ := os.Getwd()
	println(getwd)

	//Mkdir 函数可以使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误。
	//os.Mkdir()

	//MkdirAll 函数可以使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回 nil，否则返回错误。
	//权限位 perm 会应用在每一个被该函数创建的目录上。如果 path 指定了一个已经存在的目录，MkdirAll 不做任何操作并返回 nil。
	//os.MkdirAll()

	//Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。
	//RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。
	//os.Remove()
	//os.RemoveAll()

	//在 Golang 中用于执行命令的库是 os/exec，exec.Command 函数返回一个 Cmd 对象，根据不同的需求，可以将命令的执行分为三种情况
	//1. 只执行命令，不获取结果
	//2. 执行命令，并获取结果（不区分 stdout 和 stderr）
	//3. 执行命令，并获取结果（区分 stdout 和 stderr）

	//打开文件获取内容
	//os.Open()
	//open, _ := os.Open("conf.txt")

}

//http://c.biancheng.net/view/5572.html os 参考链接
//https://golang.iswbm.com/c05/c05_02.html / https://zhuanlan.zhihu.com/p/296409942 exec 参考链接
