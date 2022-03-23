package main

import (
	"fmt"
)

//猴子吃桃，第二天吃剩下的一半再多吃一个，第十天只剩1个桃子
func peach(n int) int {
	if n < 1 || n > 10 {
		fmt.Println("error days")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
func main() {
	day := 0
	fmt.Println("输入一个整数(1-10)：")
	fmt.Scanln(&day)
	fmt.Println(peach(day))

}
