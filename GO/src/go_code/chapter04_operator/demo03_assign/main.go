package main

import (
	"fmt"
)

func main() {
	//赋值运算符的使用
	//var i int
	//i = 10 基本
	//change
	a := 9
	b := 2
	fmt.Println(a, b)
	temp := a
	a = b
	b = temp
	fmt.Println(a, b)
	//复合赋值
	a += 17
	fmt.Println(a, b)

}
