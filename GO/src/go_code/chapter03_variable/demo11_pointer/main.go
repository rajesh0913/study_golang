package main

import (
	"fmt"
	"go_code/chapter03_variable/packagetest"
)

func main() {
	var i int = 10
	//取地址
	fmt.Println("i 的地址：", &i)
	//1.ptr是一个指针变量
	//2.ptr的类型 *int
	//3.ptr 本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n", ptr)
	fmt.Printf("ptr的地址：%v\n", &ptr)
	fmt.Printf("ptr指向的值：%v\n", *ptr)

	//写一个程序，获取一个int变量num的地址，并显示到终端
	//将num的地址复制给指针ptr，并通过ptr去修改num的值
	num := 99
	var ptr1 *int = &num
	fmt.Printf("num的值：\t%d\nnum的地址：\t%v\n", num, &num)
	*ptr1 = 55
	fmt.Printf("num的值：\t%v\n", num)

	num1 := packagetest.TestName
	var ptr2 *string = &num1
	fmt.Printf("num1的值：\t%s\nnum1的地址：\t%v\n", num1, &num1)
	*ptr2 = "hello a"
	fmt.Printf("num1的值：\t%v\n", num1)

}
