package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

/*
fmt包实现了类似C语言printf和scanf的格式化I/O。主要分为向外输出内容和获取输入内容两大部分。

*/

func main() {
	// // Print系列函数会将内容输出到系统的标准输出，区别在于:
	// // Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	// fmt.Print("在终端打印信息")
	// name := "fmt包学习"
	// fmt.Printf("这个是关于golang的%s\n", name)
	// fmt.Println("在终端打印单独一行显示")

	// // Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	// // 向标准输出写入内容
	// fmt.Fprint(os.Stdout, "向标准输出写入内容")
	// fileObj, err := os.OpenFile("./fmt.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("打开文件错误， err:", err)
	// 	return
	// }
	// file_content := "this is a world"
	// for i := 0; i < 10; i++ {
	// 	fmt.Fprintf(fileObj, "往文件中写如信息：%d\n", i)
	// }
	// fmt.Fprintf(fileObj, "往文件中写如信息：%s\n", file_content)

	// // Sprint系列函数会把传入的数据生成并返回一个字符串。
	// s1 := fmt.Sprint("Sprint函数")
	// age := 18
	// s2 := fmt.Sprintf("age:%d", age)
	// s3 := fmt.Sprintln("Sprintln函数")
	// fmt.Println(s1, s2, s3)

	// // Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。
	// e := errors.New("原始错误e")
	// w := fmt.Errorf("Wrap了一个错误%w", e)
	// fmt.Println(w)

	// 获取输入
	// var (
	// 	name    string
	// 	age     int
	// 	married bool
	// )
	// fmt.Scan(&name, &age, &married)
	// fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
	var command string
	fmt.Scan(&command)
	fmt.Printf("cmd is %s", command)
	cmd := exec.Command("/bin/bash", "-c", command)

	//创建获取命令输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)

}
