package main

// 值接收者实现接口
// 使用值接收者实现接口之后，不管是结构体类型还是对应的结构体指针类型的变量都可以赋值给该接口变量。

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}

// 值接收者实现接口

// Dog 结构体类型
type Dog struct {
}

// Move 使用值接收者定义Move方法实现Mover接口
func (d Dog) Move() {
	fmt.Println("dog run")
}

// // 此时实现Mover接口的是Dog类型。

func main() {
	// 声明一个Mover类型的变量x
	var x Mover
	// d1是Dog类型
	var d1 = Dog{}
	// 可以将d1赋值给变量x
	x = d1
	x.Move()

	// d2是Dog指针类型
	var d2 = &Dog{}
	x = d2
	x.Move()
}
