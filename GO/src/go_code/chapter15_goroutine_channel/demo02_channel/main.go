package main

import "fmt"

/*
	1. channel本质就是一个数据结构-队列
	2. FIFO ->先进先出
	3.线程安全，多goroutine访问不需要加锁
	4. channel有数据类型，必须一一对应
	5. channel为引用各类型，必须初始化(make)后才可使用
*/

func main() {
	//1.创建可以存放3个int类型的channel
	var intChan chan int = make(chan int, 3)

	//2.intChan的值(&intChan)
	fmt.Println(intChan)
	fmt.Println(&intChan)
	fmt.Printf("%p", &intChan)
	//3.向channel写入数据
	intChan <- 10
	num := 121
	intChan <- num
	intChan <- 23
	//	intChan <- 58

	//当给管道写入数据，不能超过管道容量
	//4.channel的长度和容量
	fmt.Printf("channel len = %v, channel cap = %v\n", len(intChan), cap(intChan))

	//5.从channel读取数据
	var num02 int = <-intChan
	fmt.Printf("num02 = %v\n", num02)
	fmt.Printf("channel len = %v, channel cap = %v\n", len(intChan), cap(intChan))
	//6.在没有使用协程的情况下，如果管道数据已经全部取除，再取会报告deadlock
	//	取出数据后可以继续存放进去
	num03 := <-intChan
	num04 := <-intChan
	//num05 := <-intChan
	fmt.Println(num02, num03, num04)

	//fmt.Println(num05)
	//7. channel关闭：只能读不能写,内置函数
	close(intChan)
	//8. channel遍历:for-range
	//	遍历时，管道没有关闭，则出现deadlock；管道已经关闭，则正常遍历，遍历完后退出
	intChan = make(chan int, 100)
	fmt.Printf("channel len = %v, channel cap = %v\n", len(intChan), cap(intChan))
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	//普通for循环不能完全遍历
	close(intChan)
	for v := range intChan {
		fmt.Println("v = ", v)
	}

}
