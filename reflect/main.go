package main

import (
	"fmt"
	"reflect"
)

// 变量的内在机制 类型信息：预先定义好的元信息。 值信息：程序运行过程中可动态变化的。
/*
* 反射是指在程序运行期对程序本身进行访问和修改的能力
* 支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。
* Go程序在运行期使用reflect包访问程序的反射信息。反射就是在运行时动态的获取一个变量的类型信息和值信息。
* 任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成
* reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
* Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
* IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
* IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic。
* IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic。
* 任意值通过reflect.TypeOf()获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象（reflect.Type）的NumField()和Field()方法获得结构体成员的详细信息。

 */

type myInt int64

func reflectType(x interface{}) {
	//v := reflect.TypeOf(x)
	//fmt.Printf("type:%v\n", v)
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。
// 而反射中使用专有的Elem()方法来获取指针对应的值。
func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		//修改的是副本，reflect包会引发panic
		v.SetInt(200)
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	//var a float32 = 3.14
	//reflectType(a) // type:float32
	//reflectValue(a)
	//var b int64 = 100
	//reflectType(b) // type:int64
	//reflectValue(b)

	//var a *float32 // 指针
	//var b myInt    // 自定义类型
	//var c rune     // 类型别名
	//reflectType(a) // type: kind:ptr
	//reflectType(b) // type:myInt kind:int64
	//reflectType(c) // type:int32 kind:int32
	//
	//type person struct {
	//	name string
	//	age  int
	//}
	//var d = person{
	//	name: "沙河小王子",
	//	age:  18,
	//}
	//type book struct{ title string }
	//var e = book{title: "《跟小王子学Go语言》"}
	//reflectType(d) // type:person kind:struct
	//reflectType(e) // type:book kind:struct

	//var a float32 = 3.14
	//var b int64 = 100
	//reflectValue(a) // type is float32, value is 3.140000
	//reflectValue(b) // type is int64, value is 100
	//// 将int类型的原始值转换为reflect.Value类型
	//c := reflect.ValueOf(10)
	//fmt.Printf("type c :%T\n", c) // type c :reflect.Value

	//var a int64 = 100
	//// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	//reflectSetValue2(&a)
	//fmt.Println(a)

	// *int类型空指针
	//var a *int
	//fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	//// nil值
	//fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	//// 实例化一个匿名结构体
	//b := struct{}{}
	//// 尝试从结构体中查找"abc"字段
	//fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())
	//// 尝试从结构体中查找"abc"方法
	//fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid())
	//// map
	//c := map[string]int{}
	//// 尝试从map中查找一个不存在的键
	//fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	stu1 := student{
		Name:  "xiaowangzi",
		Score: 10,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
}
