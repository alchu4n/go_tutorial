package main

import (
	"errors"
	"fmt"
	"strings"
)

//定义全局变量num
var num int64 = 10

// 函数是组织好的、可重复使用的、用于执行指定任务的代码块
// 函数、匿名函数和闭包
/*
	函数定义 使用func关键字,格式如下：
	func 函数名(参数)(返回值){
		函数体
	}
*/
// 求两个数之和
func intSum(x int, y int) int {
	return x + y
}

// 类型简写
func intSumSiple(x, y int) int {
	return x + y
}

// 可变参数 通过在参数名后加...来标识, 可变参数通常要作为函数的最后一个参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

// 函数的参数和返回值都是可选的
func sayHello() {
	fmt.Println("Hello 沙河")
}

// 返回值 通过return关键字向外输出,支持多返回值,函数如果有多个返回值时必须用()将所有返回值包裹起来
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 返回值命名
func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// 返回值补充 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil // 没必要返回[]int{}
// 	}
// }

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}

func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

// 定义函数类型
// 定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
type calculation func(int, int) int

// 凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 函数作为参数
func calc3(x, y int, op func(int, int) int) int {
	return op(x, y)
}

// 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc4(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func main() {
	// 变量f_closure是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f_closure的生命周期内，变量x也一直有效。
	var f_closure = adder()
	fmt.Println(f_closure(10)) //10
	fmt.Println(f_closure(20)) //30
	fmt.Println(f_closure(30)) //60

	f1_closure := adder()
	fmt.Println(f1_closure(40)) //40
	fmt.Println(f1_closure(50)) //90

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	f0, f00 := calc4(10)
	fmt.Println(f0(1), f00(2)) //11 9
	fmt.Println(f0(3), f00(4)) //12 8
	fmt.Println(f0(5), f00(6)) //13 7

	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)

	// 可变参数
	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160

	calc1, calc2 := calc(100, 20)
	fmt.Println(calc1, calc2)

	testGlobalVar() //num=10
	testLocalVar()

	// add和sub都能赋值给calculation类型的变量。
	var c calculation
	c = add
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f

	// 函数作为参数
	ret12 := calc3(10, 20, add)
	fmt.Println(ret12) //30
	d, err := do("+")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(d(1, 21))
	}

	// 匿名函数,多用于实现回调函数和闭包
	/*
	   func(参数)(返回值){
	       函数体
	   }
	*/
	// 将匿名函数保存到变量
	add_anonymous := func(x, y int) {
		fmt.Println(x + y)
	}
	// 通过变量调用匿名函数
	add_anonymous(10, 20)

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	// defer语句会将其后面跟随的语句进行延迟处理
	// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
	// 也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
	// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")

	funcA()
	funcB()
	funcC()
}
