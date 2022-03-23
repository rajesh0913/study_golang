package main

import (
	"fmt"
	"go_code/chapter06_function/demo05_init/utils"
)

func test() int {
	fmt.Println("test...")
	return 90
}

//在init()函数中完成初始化
//全局变量定义->init()函数->main()函数【重点】
func init() {

	fmt.Println("init...")

}

var age = test() //全局变量定义->init()函数->main()函数【重点】

func main() {
	fmt.Println("main...age = ", age)
	fmt.Println("Age = ", utils.Age, "Name = ", utils.Name)

}
