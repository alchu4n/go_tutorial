package main

import "fmt"

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

// 指针类型的接收者
// SetAge 设置p的年龄 使用指针接收者
/*
什么时候应该使用指针类型接收者
- 需要修改接收者中的值
- 接收者是拷贝代价比较大的大对象
- 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
*/
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// 值类型的接收者
// SetAge2 设置p的年龄
// 当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
	p1.SetAge(30)
	fmt.Println(p1.age) // 30

	p2 := NewPerson("小王子", 25)
	fmt.Println(p2.age) // 25
	p2.SetAge2(30)      // (*p1).SetAge2(30)
	fmt.Println(p2.age) // 25
}
