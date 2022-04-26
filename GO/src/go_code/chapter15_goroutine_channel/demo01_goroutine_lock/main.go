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

/*
	全局变量加锁同步解决goroutine通讯不够完美
	1.主线程是在等待所有goroutine全部完成的时间很难确定
	2.如果主线程休眠时间长了，加长等待时间；短则使goroutine处于工作状态就随着主线程退出而销毁
	3.通过全局变量加锁实现通讯，并不利于多个协程对全局变量的读写操作
*/
