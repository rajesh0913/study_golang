package main

import "fmt"

type Cat struct {
	Name string
	Age  int
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
}
