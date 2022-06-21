package main

// 单元测试 Go语言中的测试依赖go test命令、
// go test命令是一个按照一定约定和组织的测试代码的驱动程序。
// 在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分
// 不会被go build编译到最终的可执行文件中。
/**
_test.go文件中有三种类型的函数：
单元测试函数
	函数名前缀为Test-测试程序的一些逻辑行为是否正确
	每个测试函数必须导入testing包
	func TestName(t *testing.T) {//}

基准测试函数
	函数名前缀为Benchmark-测试函数的性能

示例函数。
	函数名前缀为Example-为文档提供示例文档

*/
