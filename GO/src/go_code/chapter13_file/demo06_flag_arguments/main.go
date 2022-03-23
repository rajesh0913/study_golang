package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//1. 原生方式获取参数
	//生成.exe文件进行测试
	fmt.Println("命令行参数：", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("args[%v] = %v\n", i, v)
	}
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	flag.Parse() //必须进行转换才能取到对应值

	fmt.Printf("user = %v\tpwd = %v\thost = %v\tport = %v", user, pwd, host, port)
}
