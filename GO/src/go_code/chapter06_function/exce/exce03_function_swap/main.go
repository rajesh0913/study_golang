package main

import (
	"fmt"
)

//swap函数交换两个数字
func swap(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}
func main() {
	num1 := 0
	num2 := 0
	fmt.Println("输入一个整数1：")
	fmt.Scanln(&num1)
	fmt.Println("输入一个整数2：")
	fmt.Scanln(&num2)
	fmt.Println(num1, num2)
	swap(&num1, &num2)
	fmt.Println(num1, num2)

}
