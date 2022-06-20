package main

import "fmt"

// Singer 接口
type Singer interface {
	Sing()
}

// Bird 结构体
type Bird struct {
	//name string
}

// NewBird 构造函数
func NewBird() *Bird {
	return &Bird{}
}

// Sing Bird 类型的Sing方法
func (b Bird) Sing() {
	fmt.Println("jijiji")
}

func main() {
	b1 := NewBird()
	b1.Sing()
}
