package main

import "fmt"

func main() {
	//golang 变量使用方式1
	//第一种：指定变量类型，声明不赋值，使用默认值
	//int 默认值为0
	var i int
	fmt.Println("i = ", i)

	//使用方式2：根据值自行判定变量类型（类型推导
	var num = 10.11
	fmt.Println("num = ", num)
	//使用方式3：省略var，注意 ：= 左侧的变量不应该是已经声明过的，否则编译错误
	//使用方式等价于： var name string   name = "tom"
	name := "tom"
	fmt.Println("name = ", name)

}
