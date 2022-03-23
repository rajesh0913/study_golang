package main

import (
	"fmt"
)

func main() {
	//如果运算数都是整数，那么除后去掉小数部分，保留整数部分
	//无视变量类型
	fmt.Println(10 / 4)
	var num1 float32 = 10 / 4
	fmt.Println(num1)

	//希望保留小数部分，则需要有浮点数参与运算
	var num2 float32 = 10.0 / 4
	fmt.Println(num2)
	//演示%使用 a % b = a - a / b * b
	fmt.Println("10 % 3 =", 10%3)
	// ++\--只能独立使用,只可以后置，不可以前置
	//num2 = i++(err)
	//num2 = i--(err)
	//if i++ > 0(err)
	var num3 int = 1
	num3++
	fmt.Println(num3)

}
