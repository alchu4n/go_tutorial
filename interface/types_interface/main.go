package main

import "fmt"

// 多种类型实现同一接口

type Mover interface {
	Move()
}

// Dog 结构体类型
type Dog struct {
	Name string
}

// Move 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

// Car 汽车结构体类型
type Car struct {
	Brand string
}

// Move Car类型实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s速度70迈\n", c.Brand)
}

// 一个接口的所有方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器结构体
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机结构体
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
func main() {
	var obj Mover

	obj = Dog{Name: "旺财"}
	obj.Move()

	obj = Car{Brand: "宝马"}
	obj.Move()

	var objw WashingMachine
	objw = haier{}
	objw.dry()
}
