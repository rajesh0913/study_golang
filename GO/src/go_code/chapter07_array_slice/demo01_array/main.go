package main

import (
	"fmt"
)

func main() {
	//1. 定义数组
	var hens [6]float64
	var totalWeight float64 = 0.0
	var averageWeight float64 = 0.0
	//2. 赋值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.0
	hens[4] = 2.0
	hens[5] = 50.0

	//3. 初始化方式
	//1
	var num1 [3]int = [3]int{1, 2, 3}
	//2
	var num2 = [3]int{1, 2, 3}
	//3
	var num3 = [...]int{1, 2, 3}
	//4
	var num4 = [...]int{1: 2, 0: 1, 3: 3}
	fmt.Println("num1: ", num1, "\nnum2: ", num2, "\nnum3: ", num3, "\nnum4: ", num4)

	//4. 遍历
	//方式1
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	//方式2:for-range
	for index, value := range num4 {
		fmt.Printf("%v = %v\t", index, value)

	}
	fmt.Println()
	for index := range num4 {
		fmt.Printf("num4[%v] = %v\t", index, num4[index])
	}

	fmt.Println()

	averageWeight = totalWeight / float64(len(hens))
	fmt.Printf("totalWeight = %v, averageWeight = %.2f\n", totalWeight, averageWeight)

	//数组地址
	var intArr [3]int
	fmt.Println("数组默认值： ", intArr)
	fmt.Printf("数组名：%p,数组第一个数的地址：%v,数组第二个数的地址： %v\n", &intArr, &intArr[0], &intArr[1])

}
