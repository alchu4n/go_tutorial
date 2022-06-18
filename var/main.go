package main

import "fmt"

var x = 100
var y = "小王子"

func foo() (int, string) {
	return 10, "Q1mi"
}

// Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且Go语言的变量声明后必须使用。
func main() {
	//标准声明格式 var 变量名 变量类型
	var name string
	var age int
	var isOk bool
	fmt.Println(name, age, isOk)
	//批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	//声明变量同时指定初始值 var 变量名 类型 = 表达式
	var name1 string = "小王子"
	var age1 int = 18
	fmt.Println(name1, age1)
	// 一次初始化多个变量
	var name2, age2 = "沙河娜扎", 28
	fmt.Println(name2, age2)
	//类型推导,让编译器根据变量的初始值推导出其类型
	var name3 = "小王子"
	var age3 = 18
	fmt.Println(name3, age3)
	//短变量声明
	// 在函数内部，可以使用更简略的 := 方式声明并初始化变量
	m := 10
	fmt.Println(m)
	//匿名变量
	// 使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，例如：
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
