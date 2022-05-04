package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	fmt.Println("connected = ", conn)
	//1.客户端发送单行数据，退出
	read := bufio.NewReader(os.Stdin) //标准输入
	//从终端读入一行返回

	//for循环输入
	for {

		line, err := read.ReadString('\n')
		if err != nil {
			fmt.Println("reading err = ", err)
			//return
		}
		//如果输入“exit”退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}

		//发送line
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn write err = ", err)
		}
		fmt.Printf("client send %d bytes data.\n", n)
	}
}
