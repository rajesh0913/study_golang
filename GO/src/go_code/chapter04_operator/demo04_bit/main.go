package main

import (
	"fmt"
)

func main() {
	//& * 取地址运算符的使用
	a := 100
	fmt.Printf("a address:\t\t%v\n", &a)

	var ptr *int = &a
	fmt.Printf("ptr point to value:\t%v\n", *ptr)

	//输出二进制
	fmt.Printf("a 的二进制：\t%b\n", a)
	//八进制
	b := 011
	fmt.Println("b = \t", b)
	//十六进制
	c := 0x11
	fmt.Println("c = \t", c)
	fmt.Println()
	//位运算
	fmt.Println(2 & 3) //补码：2-0000 0010 3-0000 0011
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)
	fmt.Println(-2 ^ 2)
	//原码：-2->1000 0010  反码：-2->1111 1101 补码：-2->1111 1110 2->0000 0010
	//得到补码：1111 1100 ->反码：1111 1011 -> 原码：1000 0100
	d := 1 >> 2
	f := 1 << 2
	fmt.Println("d = ", d, "f = ", f)

}
