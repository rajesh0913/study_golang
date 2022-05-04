package main

import (
	"fmt"
	"time"
)

type Cat struct {
	Name string
	Age  int
}

//4.
func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello world")
		time.Sleep(time.Millisecond * 50)
	}
}
func test() {
	//可以使用defer + recover解决
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("发生异常", err)
		}
	}()
	var mymap map[int]string = make(map[int]string, 3) //测试细节4，删除make部分
	mymap[0] = "golang"                                //error,didn't make map
}

func main() {
	//1.定义一个存放【任意】数据类型的管道存放3个数据
	Ch := make(chan interface{}, 3)
	Ch <- 10
	Ch <- "tom jack"
	Ch <- Cat{"小花猫", 4}
	//希望获取管道中第三个元素->推出前三个数据
	<-Ch
	<-Ch
	newCat := <-Ch
	fmt.Printf("newCat = %T, newCat = %v\n", newCat, newCat)

	//fmt.Printf("newCat Name = %v", newCat.Name)
	//->报错type interface {} is interface with no methods

	a := newCat.(Cat) //类型断言
	fmt.Printf("newCat Name = %v\n", a)

	//2. channel可以声明为只读或者只写，默认可读可写
	//	读写操作方法的参数会使用只读只写，防止误操作
	//	只写
	var chan1 chan<- int = make(chan<- int, 3)
	chan1 <- 20
	fmt.Println("chan1 = ", chan1)

	//只读
	//var chan2 <-chan int
	//num1 := <-chan2
	//fmt.Println("num1 = ", num1)

	//3.使用select可以解决管道阻塞的问题
	chan3 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		chan3 <- i
	}
	chan4 := make(chan string, 5)
	for j := 0; j < 5; j++ {
		chan4 <- "hello" + fmt.Sprintf("%d", j)
	}
	//传统方法遍历管道时不关闭，会导致deadlock
	//实际开发时不好确定何时关闭管道，使用select方式解决
Here:
	for {
		select {
		case v := <-chan3: //如果chan3一直没有关闭，也不会阻塞导致死锁
			fmt.Printf("从chan3读取数据：%v\n", v)
			time.Sleep(time.Millisecond * 10)
		case v := <-chan4:
			fmt.Printf("从chan4读取数据：%v\n", v)
			time.Sleep(time.Millisecond * 10)
		default:
			fmt.Printf("都取不到，退出...\n")
			time.Sleep(time.Millisecond * 10)
			break Here
		}
	}

	//4.goroutine协程中出现异常导致程序崩溃，使用defer+recover进行捕获，其他协程可以继续工作
	go sayHello()
	go test()

	for n := 0; n < 10; n++ {
		fmt.Println("ok")
		time.Sleep(time.Millisecond * 100)
	}

}
