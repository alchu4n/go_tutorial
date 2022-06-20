package main

import (
	"fmt"
)

// 空接口
/*
指没有定义任何方法的接口类型。因此任何类型都可以视为实现了空接口。
也正是因为空接口类型的这个特性，空接口类型的变量可以存储任意类型的值。
*/

// Any 不包含任何方法的空接口类型
type Any interface{}

// Dog 狗结构体
type Dog struct{}

// 使用空接口实现可以接收任意类型的函数参数
// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	var x Any
	// 字符串型
	x = "nihao"
	fmt.Printf("type:%T value:%v\n", x, x)
	x = 100 // int型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = true // 布尔型
	fmt.Printf("type:%T value:%v\n", x, x)
	x = Dog{} // 结构体类型
	fmt.Printf("type:%T value:%v\n", x, x)

	// 通常我们在使用空接口类型时不必使用type关键字声明，可以像下面的代码一样直接使用interface{}。
	var y interface{}
	y = "nihao"
	fmt.Printf("type:%T value:%v\n", y, y)

	show(x) // type:main.Dog value:{}
	show(1) // type:int value:1

	// 空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo) // map[age:18 married:false name:沙河娜扎]
}
