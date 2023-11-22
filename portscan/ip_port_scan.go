package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// 定义一个端口扫描的结构体
type PortScan struct {
	ip    string        // 要扫描的IP地址
	start int           // 起始端口
	end   int           // 结束端口
	timeout time.Duration // 超时时间
}

// 创建一个新的端口扫描对象
func NewPortScan(ip string, start, end int, timeout time.Duration) *PortScan {
	return &PortScan{
		ip: ip,
		start: start,
		end: end,
		timeout: timeout,
	}
}

// 执行端口扫描
func (ps *PortScan) Scan() {
	// 创建一个等待组，用于等待所有的协程完成
	var wg sync.WaitGroup
	// 遍历所有的端口
	for port := ps.start; port <= ps.end; port++ {
		// 为每个端口创建一个协程
		wg.Add(1)
		go func(p int) {
			// 扫描结束后，减少等待组的计数
			defer wg.Done()
			// 拼接IP地址和端口号
			address := fmt.Sprintf("%s:%d", ps.ip, p)
			// 尝试连接该地址
			conn, err := net.DialTimeout("tcp", address, ps.timeout)
			// 如果没有错误，说明端口是开放的
			if err == nil {
				// 打印IP地址和端口号
				fmt.Printf("%s:%d open\n", ps.ip, p)
				// 关闭连接
				conn.Close()
			}
		}(port)
	}
	// 等待所有的协程完成
	wg.Wait()
}

// 从文件中读取IP地址列表
func readIPs(filename string) ([]string, error) {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// 延迟关闭文件
	defer file.Close()
	// 创建一个切片，用于存储IP地址
	var ips []string
	// 创建一个缓冲读取器，用于按行读取文件内容
	scanner := bufio.NewScanner(file)
	// 循环读取每一行
	for scanner.Scan() {
		// 获取当前行的内容
		line := scanner.Text()
		// 将内容添加到切片中
		ips = append(ips, line)
	}
	// 返回切片和错误
	return ips, scanner.Err()
}

// 主函数
func main() {
	// 从文件中读取IP地址列表，文件名为ips.txt
	ips, err := readIPs("ips.txt")
	if err != nil {
		// 如果有错误，打印错误并退出程序
		fmt.Println(err)
		os.Exit(1)
	}
	// 遍历IP地址列表
	for _, ip := range ips {
		// 创建一个端口扫描对象，扫描每个IP地址的1-1024端口，超时时间为1秒
		ps := NewPortScan(ip, 1, 1024, time.Second)
		// 执行端口扫描
		ps.Scan()
	}
}
