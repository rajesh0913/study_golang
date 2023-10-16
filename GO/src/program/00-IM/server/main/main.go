package main

import (
	"GO/src/program/00-IM/server/model"
	"GO/src/program/00-IM/server/process"
	"fmt"
	"net"
	"time"
)

// 处理与客户端的通讯
func mainProcess(conn net.Conn) {
	// 延时关闭conn
	defer conn.Close()
	// 创建主控实例
	processor := &process.Processor{
		Conn: conn,
	}
	(*processor).Process2()
}

// 创建UserDao实例
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
func main() {
	// 初始化redis连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	process.InitUserMgr()
	fmt.Println("服务器在8889端口进行监听...")
	listen, err := net.Listen("tcp", "192.168.5.9:8889")
	if err != nil {
		fmt.Printf("net.listen err: %v.", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("服务器等待连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("net.listen err: %v.", err)
		}
		// 连接成功，启动协程与客户端保持通讯
		go mainProcess(conn)
	}

}
