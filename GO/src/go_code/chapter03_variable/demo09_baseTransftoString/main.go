package main

import (
	"fmt"
	"strconv"
)

//演示基本数据类型转换
func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Printf("i = %v  n1 = %v  n1 数据类型 %T\n\n", i, n1, n1)
	//基本类型转string类型（重点）
	//方式1:fmt.Sprintf("%参数"，表达式)
	var num1 int = 99
	var num2 float64 = 23.456
	var num3 bool = true
	var mtchar byte = 'h'

	//方式1
	fmt.Printf("方式1：\n")
	var str string = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	str = fmt.Sprintf("%t", num3)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	str = fmt.Sprintf("%c", mtchar)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	//方式2：使用strconv包中的函数
	fmt.Printf("方式2：\n")
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 5, 64)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

	str = strconv.FormatBool(num3)
	fmt.Printf("str type %T\t\tstr = %q\n", str, str)

}
