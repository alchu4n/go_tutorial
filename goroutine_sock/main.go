package main

import (
	"fmt"
	"sync"
)

// 并发安全和锁

var (
	x  int64
	wg sync.WaitGroup
	// 互斥锁 sync.Mutex
	m sync.Mutex
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		// 修改x前加锁
		m.Lock()
		x = x + 1
		// 修改完后解锁
		m.Unlock()
	}
	wg.Done()
}

// 代码编译后执行，不出意外每次执行都会输出诸如9537、5865、6527等不同的结果
// 开启了两个 goroutine 分别执行 add 函数，这两个 goroutine 在访问和修改全局的x变量时就会存在数据竞争
// 某个 goroutine 中对全局变量x的修改可能会覆盖掉另一个 goroutine 中的操作，所以导致最后的结果与预期不符。

// 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源。
// Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
