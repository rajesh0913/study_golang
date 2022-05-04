package main

import (
	"fmt"
	"time"
)

/* 协同工作案例
1.开启writeData协程，向管道intChan写入50个整数
2.开启一个readData协程，从管道intChan读取写入的数据
3.writeData和readData操作同一个管道
4.主线程需要等待协程都完成工作才能退出
*/
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData写入数据：%v\n", i)
		time.Sleep(time.Millisecond * 1)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData读到数据：%v\n", v)
		time.Sleep(time.Millisecond * 1)
	}
	exitChan <- true
	close(exitChan)
}

func main() {

	inChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(inChan)
	go readData(inChan, exitChan)
	//time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
