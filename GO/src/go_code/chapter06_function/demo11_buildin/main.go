package main

import (
	"fmt"
)

func main() {

	//len

	//new:分配内存，主要用来分配值类型，如int、float,返回指针
	num1 := 100
	fmt.Printf("num1: type = %T\tvalue = %v\taddress = %v\n", num1, num1, &num1)

	num2 := new(int)
	fmt.Printf("num2: type = %T\tvalue = %v\taddress = %v\t*num2 = %v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2: type = %T\tvalue = %v\taddress = %v\t*num2 = %v\n", num2, num2, &num2, *num2)

	//make:分配内存，主要用来分配引用类型，如chan、map、slice

}
