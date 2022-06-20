package main

import "fmt"

// 一个接口类型的变量能够存储所有实现了该接口的类型变量

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("汪汪汪~")
}

type Sayer interface {
	Say()
}

func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	//var c Cat
	//MakeHungry(c)
	//var d Dog
	//MakeHungry(d)
	var s Sayer
	a := Cat{}
	d := Dog{}
	s = a
	s.Say()
	s = d
	s.Say()

}
