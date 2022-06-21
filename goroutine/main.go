package main

import (
	"fmt"
	"sync"
)

/*
* 进程（process）：程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
* 线程（thread）：操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位。
* 协程（coroutine）：非操作系统提供而是由用户自行创建和控制的用户态‘线程’，比线程更轻量级。
 */

// 声明全局等待组变量
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	wg.Done() // 告知当前goroutine完成
}

// 启动多个goroutine
func hello1(i int) {
	// goroutine结束就登记-1
	defer wg.Done()
	fmt.Println("hello", i)
}

// 在Go 程序启动时, Go 程序就会为 main 函数创建一个默认的 goroutine
func main() {
	// 登记1个goroutine
	//wg.Add(1)
	// 启动 goroutine 的方式非常简单，只需要在调用函数（普通函数和匿名函数）前加上一个go关键字。
	//go hello()
	//fmt.Println("你好")
	// 因为在程序中创建 goroutine 执行函数需要一定的开销，而与此同时 main 函数所在的 goroutine 是继续执行的。
	//time.Sleep(time.Second)
	//wg.Wait() // 阻塞等待登记的goroutine完成

	// 多次执行上面的代码会发现每次终端上打印数字的顺序都不一致。这是因为10个 goroutine 是并发执行的，而 goroutine 的调度是随机的。
	for i := 0; i < 10; i++ {
		// 启动一个goroutine就登记+1
		wg.Add(i)
		go hello1(i)
	}
	// 等待所有登记的goroutine都结束
	wg.Wait()

}
