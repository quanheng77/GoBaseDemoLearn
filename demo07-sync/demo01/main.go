package main

import (
	"fmt"
	"sync"
)

func main() {
	for1000()
}

func for1000() {
	group := sync.WaitGroup{}
	group.Add(100)
	var num int
	for i := 0; i < 100; i++ {
		num++
		go func(num int) {
			fmt.Println(num)
			defer group.Done()
		}(num)

	}
	group.Wait()
}

/*
主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep() 来睡眠一段时间，
等待其他线程充分运行。对于简单的代码，100个for循环可以在1秒之内运行完毕，time.Sleep() 也可以达到想要的效果。

但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。
这时候就不能使用time.Sleep() 来完成等待操作了。

WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait()
用来控制计数器的数量。Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。


WaitGroup对象不是一个引用类型，在通过函数传值的时候需要使用地址：
*/
