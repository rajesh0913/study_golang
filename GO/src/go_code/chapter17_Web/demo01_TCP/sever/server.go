package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端数据
	defer conn.Close() //关闭conn
	for {
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.没有发送，则协程阻塞
		//fmt.Println("server waiting..." + conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("servers read err = ", err)
			return
		}
		//3.显示客户端发送内容
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听...")
	//1.tcp表示使用网络协议为tcp
	//2.在本地监听，端口为8888
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}

	defer listen.Close() //延时关闭

	//循环等待客户端链接
	for {
		fmt.Println("waiting client connecting...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err = ", err)
			//return
		} else {
			fmt.Printf("Accept() success conn = %v, client ip = %v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
	//fmt.Printf("listen success = %v", listen)
}
