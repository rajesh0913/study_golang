package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	//将接口变量(一般类型)赋值给自定义类型的变量(指定类型)->类型断言
	var b Point
	var isOk bool
	b, isOk = a.(Point)
	//检查机制
	if isOk {
		fmt.Println(b, "Convert success!")
	} else {
		fmt.Println("Convert fail!")
	}
	fmt.Println("Progress continue")
}
