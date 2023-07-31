package main

import (
	"fmt"
	"go_code/chapter10_object/demo05_factory/model"
)

func main() {
	//工厂模式使用样例，跨包使用非全局作用域结构体
	Stu := model.NewStudent("tom", 22)
	fmt.Println(*Stu)
	//Name-Age 可以直接调用
	fmt.Println(Stu.Name, Stu.Age)
	fmt.Printf("%p\n", Stu)
	//score不能直接调用,通过特定函数取值
	fmt.Println("score = ", Stu.GetScore())

}
