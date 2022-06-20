package main

import "fmt"

// 接口值
/*
由于接口类型的值可以是任意一个实现了该接口的类型值，所以接口值除了需要记录具体值之外，还需要记录这个值属于的类型。
也就是说接口值由“类型”和“值”组成，鉴于这两部分会根据存入值的不同而发生变化，我们称之为接口的动态类型和动态值。
*/

type Mover interface {
	Move()
}

type Dog struct {
	Name string
}

type Car struct {
	Brand string
}

func (d *Dog) Move() {
	println(d.Name + " Dog running")
}

func (c *Car) Move() {
	println("Car running")
}

// 如果对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现。

// justifyType 对传入的空接口类型变量x进行类型断言
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	// 接口变量m是接口类型的零值，也就是它的类型和值部分都是nil
	// type nil ;  value nil
	var m Mover
	// 使用m == nil来判断此时的接口值是否为空
	fmt.Println(m == nil) // true
	// 我们不能对一个空接口值调用任何方法，否则会产生panic。
	// m.Move() // panic: runtime error: invalid memory address or nil pointer dereference
	// 将一个*Dog结构体指针赋值给变量m
	// 接口值m的动态类型会被设置为*Dog，动态值为结构体变量的拷贝。
	m = &Dog{Name: "douban"}
	m.Move()
	var c *Car
	// 接口值m的动态类型为*Car，动态值为nil。
	m = c

	var n Mover = &Dog{Name: "旺财"}
	v, ok := n.(*Dog)
	if ok {
		fmt.Println("类型断言成功")
		v.Name = "富贵" // 变量v是*Dog类型
	} else {
		fmt.Println("类型断言失败")
	}

	justifyType(n)
}
