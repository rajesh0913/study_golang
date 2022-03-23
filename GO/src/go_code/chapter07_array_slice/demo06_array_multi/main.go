package main

import "fmt"

func main() {
	//定义1
	var arr [4][6]int
	//赋值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	//定义2
	brr := [4][6]int{}
	//遍历1：常规for循环
	for j := 0; j < 4; j++ {
		fmt.Println(brr[j])
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ { //每一列的长度也可以用len获取
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}
	//遍历2：for-range
	for index, value := range arr {
		fmt.Printf("%v-%v\n", index, value)
	}
	for index1, value1 := range arr {
		for index2, value2 := range value1 {
			fmt.Printf("arr[%v][%v] = %v\t", index1, index2, value2)
		}
		fmt.Println()
	}
}
