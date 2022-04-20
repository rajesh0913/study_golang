package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1-200各个数的累加，并且存入map中
//1.函数计算累加，放入map
//2.启动多个协程，统计结果放入map

var (
	myMap = make(map[int]int, 10)
	//lock 是一个全局的互斥锁
	lock sync.Mutex
)

//1
func test(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	//休眠2S
	time.Sleep(time.Second * 2)

	//设计层面休眠2秒，没有冲突，但是底层不知道
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%v]=%v\n", i, v)
	}
	fmt.Println(myMap[200])
	lock.Unlock()
}
