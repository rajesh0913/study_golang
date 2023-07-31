package main

import (
	"errors"
	"fmt"
	"time"
)

func test() {
	//defer + recover捕捉异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err!!!")
		}
	}()
	n1 := 10
	n2 := 0
	fmt.Println("n1/n2 = ", n1/n2)
}

//读取配置文件
//读取失败返回自定义错误
func file(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取配置文件失败")
	}
}
func test02(str string) {
	err := file(str)
	if err != nil {
		panic(err)
	}
	fmt.Println("test02函数下面的代码...")
}

func main() {
	//错误机制2
	//支持自定义错误
	//errors.new("错误说明")，返回一个error类型的值，表示错误
	//panic内置函数，接收一个interfac类型的值，可以接收error，【输出错误信息，终止程序】
	test02("config.ini")

	//错误处理机制1
	//1. 发生错误，捕捉错误，处理保证程序继续运行
	//2. 给管理员提示
	//go中抛出panic异常，defer中通过recover捕获异常，正常处理
	test() //panic
	for {

		fmt.Println("main函数下面的代码...")
		time.Sleep(time.Second)
	}

}
