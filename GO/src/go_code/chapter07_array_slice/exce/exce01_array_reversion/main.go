package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成5个数，反转打印
func main() {
	rand.Seed(time.Now().Unix())
	num := [5]int{}
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100) //0<=num[i]<100

	}
	fmt.Println("num[5] = ", num)

	for i := len(num) - 1; i >= 0; i-- {
		fmt.Println(num[i])
	}
	//反转数组
	temp := 0
	for i := 0; i < len(num)/2; i++ {
		temp = num[len(num)-i-1]
		num[len(num)-i-1] = num[i]
		num[i] = temp
	}
	fmt.Println(num)

}
