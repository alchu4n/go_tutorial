package main

import "fmt"

type Cat struct {
}

func (c Cat) Say() {
	fmt.Println("喵喵喵")
}

type Dog struct {
}

func (d Dog) Say() {
	fmt.Println("汪汪汪")
}

type Sheep struct{}

func (s Sheep) Say() {
	fmt.Println("咩咩咩~")
}

// 接下来定义一个饿肚子的场景

// MakeCatHungry 猫饿了会喵喵喵
func MakeCatHungry(c Cat) {
	c.Say()
}

// MakeSheepHungry 羊饿了会咩咩咩~
func MakeSheepHungry(s Sheep) {
	s.Say()
}

//.....

// 在饿肚子这个场景下，我们可不可以把所有动物都当成一个“会叫的类型”来处理呢？
// 当然可以！使用接口类型就可以实现这个目标。 我们的代码其实并不关心究竟是什么动物在叫，
// 我们只是在代码中调用它的Say()方法，这就足够了

type Sayer interface {
	Say()
}

// MakeHungry 饿肚子了...
func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	c := Cat{}
	c.Say()
	d := Dog{}
	d.Say()
	c1 := Cat{}
	MakeHungry(c1)
	d2 := Dog{}
	MakeHungry(d2)

}
