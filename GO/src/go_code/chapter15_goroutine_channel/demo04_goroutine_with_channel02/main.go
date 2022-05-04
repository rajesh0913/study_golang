package main

import (
	"fmt"
)

/* 协同工作案例：统计1-20000中的素数
1.开启putNum协程，向管道intChan写入整数
2.开启四个primeNum协程，从管道intChan读取写入的数据并计算是否为素数
3.素数存入primeChan
4.主线程需要等待协程都完成工作才能退出
*/
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
		//fmt.Printf("putNum写入数据：%v\n", i)
		//time.Sleep(time.Millisecond * 1)
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
		// fmt.Printf("readData读到数据：%v\n", num)
		// time.Sleep(time.Millisecond * 1)
	}
	fmt.Printf("有一个协程取不到数据\n")
	exitChan <- true
	//close(exitChan)
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 4)

	//1.开启一个写入整数的协程
	go putNum(intChan)
	//2.开启四个读取和判断的协程
	for i := 1; i <= 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	//3.工作结束判断
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(primeChan)
	}()

	//4.统计素数
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}
	//5. 结束
	fmt.Println("主线程退出...")
}
