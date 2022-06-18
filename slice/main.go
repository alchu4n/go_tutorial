package main

import "fmt"

// 数组局限性-长度固定。
// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
func main() {
	// 声明切片类型的基本语法如下 var name []T
	// var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	// fmt.Println(a)        //[]
	fmt.Println(b) //[]
	fmt.Println(c) //[false true]
	// fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较

	// 切片的长度和容量 len() cap()
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	/*
		a[2:]  // 等同于 a[2:len(a)]
		a[:3]  // 等同于 a[0:3]
		a[:]   // 等同于 a[0:len(a)]
	*/
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// 完整切片表达式 a[low : high : max]
	t := a[1:3:5]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	// 使用make()函数构造切片 make([]T, size, cap)
	d := make([]int, 2, 10)
	fmt.Println(d)      //[0 0]
	fmt.Println(len(d)) //2
	fmt.Println(cap(d)) //10

	// 判断切片是否为空 请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
	var z1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	z2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	z3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
	fmt.Printf("z1:%v len(z1):%v cap(z1):%v\n", z1, len(z1), cap(z1))
	fmt.Printf("z2:%v len(z2):%v cap(z2):%v\n", z2, len(z2), cap(z2))
	fmt.Printf("z3:%v len(z3):%v cap(z3):%v\n", z3, len(z3), cap(z3))

	// 切片的赋值拷贝 拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	s3 := make([]int, 3) //[0 0 0]
	s4 := s3             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s4[0] = 100
	fmt.Println(s3) //[100 0 0]
	fmt.Println(s4) //[100 0 0]

	// 切片遍历
	s5 := []int{1, 3, 5}

	for i := 0; i < len(s5); i++ {
		fmt.Println(i, s[i])
	}

	for index, value := range s5 {
		fmt.Println(index, value)
	}

	// 添加元素
	var s6 []int
	s6 = append(s6, 1)       // [1]
	s6 = append(s6, 2, 3, 4) // [1 2 3 4]
	s7 := []int{5, 6, 7}
	s6 = append(s6, s7...) // [1 2 3 4 5 6 7]

	// 使用copy()函数复制切片 copy(destSlice, srcSlice []T)
	// copy()复制切片
	a3 := []int{1, 2, 3, 4, 5}
	c3 := make([]int, 5, 5)
	copy(c3, a3)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a3) //[1 2 3 4 5]
	fmt.Println(c3) //[1 2 3 4 5]
	c3[0] = 1000
	fmt.Println(a3) //[1 2 3 4 5]
	fmt.Println(c3) //[1000 2 3 4 5]

}
