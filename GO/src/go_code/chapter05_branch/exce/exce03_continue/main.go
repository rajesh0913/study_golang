package main

import (
	"fmt"
)

func main() {
	//接收一个整数，打印对应层数的金字塔
	num := 0

	fmt.Println("输入金字塔层数：")
	fmt.Scanln(&num)

	//打印矩形
	for i := 0; i < num; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印半个金字塔
	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印整个金字塔
	for i := 1; i <= num; i++ {
		for m := 1; m <= num-i; m++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//打印空心金字塔
	for i := 1; i <= num; i++ {
		for m := 1; m <= num-i; m++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == num {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
