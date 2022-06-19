package main

import (
	"fmt"
	"unsafe"
)

// Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。
// Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。

// 基本的数据类型 如string、整型、浮点型、布尔等数据类型
// 自定义类型 使用type关键字来定义自定义类型。

// 自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义
//类型定义
type NewInt int

//类型别名
type MyInt = int

// 自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体 使用用ype和struct关键字来定义结构体
/*
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
*/

// 定义结构体
type person struct {
	name string
	city string
	age  int8
}

// 构造函数实现 struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// 方法和接收者
// Go语言中的方法（Method）是一种作用于特定类型变量的函数。
// 这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
/*
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
		例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
方法名、参数列表、返回参数：具体格式与函数定义相同。
*/
//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {

	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	// 结构体事例化 var 结构体实例 结构体类型
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	// 匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	// 通过使用new关键字对结构体进行实例化，得到的是结构体的地址 --> 指针类型结构体
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person 可以看出p2是一个结构体指针。
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	// 在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}

	// 取结构体的地址实例化
	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p3 := &person{}            // 等价于 var p2 = new(person)
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	// 结构体初始化-没有初始化的结构体，其成员变量都是对应其类型的零值。
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
	// 使用键值对初始化
	// 使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

	// 对结构体指针进行键值对初始化
	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}

	// 当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。
	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}

	// 使用值的列表初始化
	// 初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：
	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}

	// 结构体内存布局 - 结构体占用一块连续的内存。
	type test struct {
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)

	// 空结构体是不占用空间的。
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

	// 调用构造函数
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}

	// 方法和接收者 方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
	p10 := NewPerson("小王子", 25)
	p10.Dream()
}
