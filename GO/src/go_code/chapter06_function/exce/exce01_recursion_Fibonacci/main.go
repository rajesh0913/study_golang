package main

import (
	"fmt"
)

//斐波那契数列
func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
func main() {
	n := 0
	fmt.Println("输入一个整数：")
	fmt.Scanln(&n)
	fmt.Println(Fibonacci(n))

}
