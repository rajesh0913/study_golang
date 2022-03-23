package main

import (
	"fmt"
)

//方式3：将匿名函数赋值给全局变量
var (
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	//方式1：定义匿名函数时就直接调用，只可调用一次
	sum := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println(sum)

	//方式2：将匿名函数赋值给一个变量
	sub := func(n1 int, n2 int) int {
		return n1 - n2
	}
	a := sub(10, 30)
	fmt.Println(a)
	a = sub(0, a)
	fmt.Println(a)
	fmt.Println(Fun1(10, 20))

}
