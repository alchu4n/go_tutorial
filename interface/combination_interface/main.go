package main

// 接口组合 接口与接口之间可以通过互相嵌套形成新的接口类型
// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}

// 对于这种由多个接口类型组合形成的新接口类型，同样只需要实现新接口类型中规定的所有方法就算实现了该接口类型。

// 接口也可以作为结构体的一个字段
// src/sort/sort.go

// Interface 定义通过索引对元素排序的接口类型
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// reverse 结构体中嵌入了Interface接口
type reverse struct {
	Interface
}

// 通过在结构体中嵌入一个接口类型，从而让该结构体类型实现了该接口类型，并且还可以改写该接口的方法。

// Less 为reverse类型添加Less方法，重写原Interface接口类型的Less方法
func (r reverse) Less(i, j int) bool {
	// 通过将索引参数交换位置实现反转
	return r.Interface.Less(j, i)
}
