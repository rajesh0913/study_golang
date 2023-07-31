package main

import "fmt"

func main() {
	/* 数据结构
	type slice struct{
		ptr* []int
		len int
		cap int
	}
	*/
	//切片是数组的引用，引用传递
	var intArr [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	var intArr2 [3]int = [3]int{1, 2, 3}
	//动态长度

	//定义
	var Slice []int = intArr2[1:2]
	fmt.Println("Slice capability = ", cap(Slice))
	//方式1：引用数组
	Slice = intArr[1:3] //左闭右开
	fmt.Println("intArr = ", intArr)
	fmt.Println("Slice = ", Slice)
	fmt.Println("Slice size = ", len(Slice))
	fmt.Println("Slice capability = ", cap(Slice))

	fmt.Printf("intArr[1] address: %p  Slice[0] address: %p\n", &intArr[1], &Slice[0])

	//修改
	Slice[1] = 44
	fmt.Println("Slice[1] = ", Slice[1])

	//方式2：make创建切片
	//make方式切片，只能通过slice进行访问，没有数组指定
	var slice2 []int = make([]int, 4, 10) //make:分配并初始化一个类型为切片、映射或者通道的对象
	slice2[0] = 10
	slice2[1] = 20
	fmt.Println(slice2)
	fmt.Println("slice2 size = ", len(slice2))
	fmt.Println("slice2 capability = ", cap(slice2))

	//方式3:定义切片， 直接指定数组
	var slice3 []int = []int{1, 2, 3}
	fmt.Println(slice3)

	//切片遍历
	//1: 常规for遍历
	for i := 0; i < len(slice3); i++ {
		fmt.Print(slice3[i], "\t")
	}
	fmt.Println()
	//2: for-range
	for index, value := range slice3 {
		fmt.Printf("slice[%v] = %v\t", index, value)
	}
	fmt.Println()
}
