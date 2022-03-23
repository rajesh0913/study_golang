package main

import (
	"fmt"
)

//支持对函数返回值命名（减少返回值顺序带来的错误
func sequence(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func calculate(n1 float64, n2 float64, operator string) float64 {
	switch operator {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	default:
		return 0
	}
}

//函数调用基础类型都是值传递，可以取地址，在函数进行修改
func test1(n *int) {
	*n = *n - 10

}

//函数作为一种数据类型也可以作为形参调用
func test4(n1 int, n2 int) int {
	return n1 + n2
}

//自定义数据类型
type Myfunc func(int, int) int

func test3(funvar Myfunc, n1 int, n2 int) int {
	return funvar(n1, n2)
}

func main() {
	//自定义数据类型
	//type MyInt int

	var num1 float64 = 10.0
	var num2 float64 = 20.0
	num3 := 10

	//通过函数修改函数作用域外的值
	test1(&num3)

	//函数也是一种数据类型,可以赋值给一个变量
	//则可以通过变量调用该函数。该变量为函数类型的变量
	style := test4
	fmt.Printf("num3 = %v    style type %T\n", num3, style)

	//函数作为一种数据类型也可以作为形参调用
	fmt.Printf("test3()传入函数形参后的结果：%v\n", test3(style, 10, 20))

	var operation string
	fmt.Println("输入数1：")
	fmt.Scanln(&num1)
	fmt.Println("输入数2：")
	fmt.Scanln(&num2)
	fmt.Println("输入操作：")
	fmt.Scanln(&operation)

	fmt.Printf("%v %v %v = %f", num1, operation, num2, calculate(num1, num2, operation))

	_, _ = sequence(1, 2)
}
