package main

import (
	"fmt"
	"strconv"
)

//演示基本数据类型转换
func main() {
	var str1 string = "true"
	var b bool
	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type\t \t%T\t\tb = %v\n", b, b)

	var str2 string = "123456"
	var num1 int64
	num1, _ = strconv.ParseInt(str2, 10, 64)
	var num2 int = int(num1)
	fmt.Printf("num1 type \t%T\t\tnum1 = %v\n", num1, num1)
	fmt.Printf("num2 type \t%T\t\tnum2 = %v\n", num2, num2)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type \t%T\t\tf1 = %v\n", f1, f1)

	var str4 string = "hello"
	var num3 int64
	num3, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("num3 type \t%T\t\tnum3 = %v\n", num3, num3)
	var b2 bool
	b2, _ = strconv.ParseBool(str4)
	fmt.Printf("b2 type \t%T\t\tb2 = %v\n", b2, b2)

}
