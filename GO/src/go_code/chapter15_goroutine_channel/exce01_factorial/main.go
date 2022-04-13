package main

import (
	"fmt"
	"time"
)

//计算1-200各个数的阶乘，并且存入map中
//1.函数计算阶乘，放入map
//2.启动多个协程，统计结果放入map

var (
	myMap = make(map[int]int, 10)
)

//1
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//休眠10S
	time.Sleep(time.Second * 10)
	for i, v := range myMap {
		fmt.Printf("map[%v]=%v", i, v)
	}
}
