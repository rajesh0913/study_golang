package main

import (
	"fmt"
)

func test(arr [3]int) { //[4]int无法传递进来
	arr[0] = 99
}

func test2(arr *[3]int) {
	(*arr)[0] = 99 //重要
}
func main() {

	//数组定义后，数据类型和长度为固定，不能动态变化（包括整型存浮点型也报错

	//中括号中不填长度为slice切片
	var slice []int
	fmt.Println(slice)

	//数组中数据类型任意，可以为值类型也可以为引用类型

	//下标从0开始 -。-

	//数组为值类型，默认情况进行值传递和值拷贝，不同数据长度也不能完成值传递，长度是数组的一部分
	arr := [3]int{1, 2, 3}
	test(arr)
	fmt.Println(arr)
	test2(&arr)
	fmt.Println(arr)

}
