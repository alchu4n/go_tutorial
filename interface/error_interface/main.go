package main

import (
	"errors"
	"fmt"
	"os"
)

// Error接口和错误处理
/*
Go 语言中的错误处理与其他语言不太一样，它把错误当成一种值来处理，更强调判断错误、处理错误，而不是一股脑的 catch 捕获异常。
不支持其他语言中使用try/catch捕获异常的方式。
*/

// Go 语言中使用一个名为 error 接口来表示错误类型
type error interface {
	// Error error 接口只包含一个方法——Error，这个函数需要返回一个描述错误信息的字符串
	Error() struct{}
}

// 当一个函数或方法需要返回错误时，我们通常是把错误作为最后一个返回值。

//Open 例如下面标准库 os 中打开文件的函数
//func Open(name string) (*File, error) {
//	return OpenFile(name, O_RDONLY, 0)
//}

// 自定义 error

// New 使用errors 包提供的New函数创建一个错误
// func New(test string) error

//它接收一个字符串参数返回包含该字符串的错误。我们可以在函数返回时快速创建一个错误。
//func queryById(id int64) (*Info, eror) {
//	if id < 0 {
//		return nil, errors.New("无效的id")
//	}
//	fmt.Println(id)
//}

// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}

func main() {
	file, err := os.Open("./xx.go")
	if err != nil {
		fmt.Println("打开文件失败,err:", err)
		// 需要传入格式化的错误描述信息时，使用fmt.Errorf
		// 这种方式会丢失原有的错误类型，只拿到错误描述的文本信息
		//fmt.Errorf("打开文件失败，err:%v", err)
		// 为了不丢失函数调用的错误链，使用fmt.Errorf时搭配使用特殊的格式化动词%w，可以实现基于已有的错误再包装得到一个新的错误。
		// fmt.Errorf("打开文件失败，err:%w", err)
		// 对于这种二次包装的错误，errors包中提供了以下三个方法。
		/*
			func Unwrap(err error) error                 // 获得err包含下一层错误
			func Is(err, target error) bool              // 判断err是否包含target
			func As(err error, target interface{}) bool  // 判断err是否为target类型
		*/
		return
	}
	// 当我们使用fmt包打印错误时会自动调用 error 类型的 Error 方法，也就是会打印出错误的描述信息。
	fmt.Println(file)

	// 标准库io.EOF错误定义
	var EOF = errors.New("EOF")
	println(EOF)

}
