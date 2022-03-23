package main

import (
	"fmt"
	"go_code/chapter06_function/demo02_package/utils"
)

func main() {
	var num1 float64 = 10.0
	var num2 float64 = 20.0
	var operation byte = '*'
	// fmt.Println("输入数1：")
	// fmt.Scanln(&num1)
	// fmt.Println("输入数2：")
	// fmt.Scanln(&num2)
	// fmt.Println("输入操作：")
	// fmt.Scanln(&operation)
	//fmt.Printf("%c", operation)
	fmt.Printf("%v %c %v = %.2f", num1, operation, num2, utils.Cal(num1, num2, operation))
}
