package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

// Go语言中的指针不能进行偏移和运算，是安全指针
// 重要的概念指针地址、指针类型和指针取值
func main() {
	// 任何程序数据载入内存后，在内存都有他们的地址，这就是指针
	// 而为了保存一个数据在内存中的地址，我们就需要指针变量。
	// var_zym_a := "永远不要高估自己"
	// var_pointer_b := "0x123456" // 指针变量
	// &（取地址） ptr := &v    // v的类型为T v:代表被取地址的变量，类型为T
	// ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
	// *（根据地址取值）
	// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	a := 10
	b := &a                            // 取变量a的地址，将指针保存到b中                           //
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018

	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	/*
		对变量进行取地址（&）操作，可以获得这个变量的指针变量。
		指针变量的值是指针地址。
		对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
	*/
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100

	// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	// 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存
	// new func new(Type) *Type
	// Type表示类型，new函数只接受一个参数，这个参数是一个类型
	// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针
	a1 := new(int)
	b1 := new(bool)
	fmt.Printf("%T\n", a1) // *int
	fmt.Printf("%T\n", b1) // *bool
	fmt.Println(*a1)       // 0
	fmt.Println(*b1)       // false
	// 声明了一个指针变量a但是没有初始化
	var a2 *int
	// 指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	// 使用内置的new函数对a进行初始化之后就可以正常对其赋值
	a2 = new(int)
	*a2 = 10
	fmt.Println(*a2)

	// make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了
	// func make(t Type, size ...IntegerType) Type
	// 在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作

	// 只是声明变量b是一个map类型的变量
	var b2 map[string]int
	// 使用make函数进行初始化操作
	b2 = make(map[string]int, 10)
	b2["沙河娜扎"] = 100
	fmt.Println(b)
}
