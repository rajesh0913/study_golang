package main

import (
	"fmt"
	"strings"
)

//闭包：一个函数与其相关的引用环境组合的一个整体
//返回类型：func(int) int
//返回值：匿名函数func(x int) int
//返回匿名函数，但是匿名函数引用到匿名函数外的n，n与匿名函数构成闭包
//n只被初始化一次，而后再调用f()，反复调用，n就会被累计
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

//编写函数，接收一个文件后缀名，返回一个闭包
//调用闭包，可以传入一个文件名，如果该文件名没有指定后缀，则返回文件名.jpg
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		} else {
			return name
		}

	}
}
func main() {

	f := AddUpper()
	fmt.Println(f(1))

	s := makeSuffix(".jpg")
	fmt.Println(s("qinzhibao"))
	fmt.Println(s("qinzhibao.jpg"))

}
