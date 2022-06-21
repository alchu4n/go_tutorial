package main

import "fmt"

/*
* 如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制
* Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
* Go 语言中的通道（channel）是一种特殊的类型。
* 通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
* 每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
 */

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道

// 使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
// 同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。
// 无缓冲的通道必须有至少一个接收方才能发送成功
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功 ", ret)
}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		} else {
			fmt.Printf("v:%#v ok:%#v\n", v, ok)
		}
	}
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 1000; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// 单向通道
//<- chan int // 只接收通道，只能接收不能发送
//chan <- int // 只发送通道，只能发送不能接收

func main() {
	// 未初始化的通道类型变量其默认零值是nil
	fmt.Println(ch1) // <nil>
	fmt.Println(ch2) // <nil>
	fmt.Println(ch3) // <nil>

	// 初始化一个通道
	ch4 := make(chan int)
	ch5 := make(chan bool, 1) // 声明一个缓冲区大小为1的通道
	fmt.Println(ch4)
	fmt.Println(ch5)

	// 通道操作 （无缓冲区的通道又称同步通道） 存在死锁问题
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	// 将一个值发送到通道中。
	ch <- 10 // 代码会阻塞在ch <- 10这一行代码形成死锁, 需要创建一个 goroutine 去接收值

	fmt.Println("发送成功")

	// 从一个通道中接收值。
	// x := <-ch
	// 从ch中接收值，忽略结果 <-ch
	// fmt.Println(x)
	//关闭通道
	//close(ch)

	// 有缓冲的通道
	ch2 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	var i int = 10
	//fmt.Print("请输入i的值：")
	//fmt.Scanln(&i)
	ch2 <- i
	fmt.Println("发送成功", i)
	y := <-ch2
	fmt.Println("接收成功", y)

	// 对一个通道执行接收操作时支持使用如下多返回值模式。
	// 判断一个通道是否被关闭
	//value, ok := <-ch2
	// value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
	// ok：通道ch关闭时返回 false，否则返回 true。
	ch6 := make(chan int, 2)
	ch6 <- 1
	ch6 <- 2
	close(ch6)
	f2(ch6)

	// for range接收值
	// func f3(){}

	// 单向通道
	//<- chan int // 只接收通道，只能接收不能发送
	//chan <- int // 只发送通道，只能发送不能接收
	ch7 := Producer()
	res := Consumer(ch7)
	fmt.Println(res)

	// select多路复用
	// Go 语言内置了select关键字，使用它可以同时响应多个通道的操作。
	/*
		可处理一个或多个 channel 的发送/接收操作。
		如果多个 case 同时满足，select 会随机选择一个执行。
		对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
	*/

	//创建了一个缓冲区大小为1的通道 ch
	ch8 := make(chan int, 1)
	/*
		第一次循环时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
		第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；
		后续的 for 循环以此类推会依次打印出3、5、7、9
	*/
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch8:
			fmt.Println(x)
		case ch8 <- i:
		}
	}
}
