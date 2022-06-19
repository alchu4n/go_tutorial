package main

import "fmt"

// 嵌套匿名字段
// 当访问结构体成员时会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

//Address2 地址结构体
type Address2 struct {
	Province string
	City     string
}

//User2 用户结构体
type User2 struct {
	Name    string
	Gender  string
	Address //匿名字段
}

// 嵌套结构体的字段名冲突

//Address3 地址结构体
type Address3 struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User3 用户结构体
type User3 struct {
	Name   string
	Gender string
	Address3
	Email
}

// 结构体的“继承”

//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	var user2 User2
	user2.Name = "小王子"
	user2.Gender = "男"
	user2.Address.Province = "山东"    // 匿名字段默认使用类型名作为字段名
	user2.City = "威海"                // 匿名字段可以省略
	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"小王子", Gender:"男", Address:main.Address{Province:"山东", City:"威海"}}

	var user3 User3
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address3.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"    //指定Email结构体中的CreateTime

	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
