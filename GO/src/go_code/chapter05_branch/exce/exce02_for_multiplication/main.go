package main

import (
	"fmt"
)

func main() {
	//接收一个整数，打印乘法表
	num := 0
	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num)
	//打印一行
	for i := 1; i <= num; i++ {
		fmt.Printf("%v * %v = %v\t", i, num, i*num)
	}
	fmt.Println()
	//打印从1开始
	for j := 1; j <= num; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%v * %v = %v\t", j, i, i*j)
		}
		fmt.Println()
	}

}
