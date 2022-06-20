package main

import "fmt"

// 一个类型实现多个接口

// Sayer 接口
type Sayer interface {
	Say()
}

// Mover 接口
type Mover interface {
	Move()
}
type Dog struct {
	Name string
}

// Say 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("%s会叫汪汪汪\n", d.Name)
}

// Move 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("%s会动\n", d.Name)
}

func main() {
	//	同一个类型实现不同的接口互相不影响使用。
	var d = Dog{Name: "旺财"}
	var s Sayer = d
	var m Mover = d
	s.Say()
	m.Move()
}
