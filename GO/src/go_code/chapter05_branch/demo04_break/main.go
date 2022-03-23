package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//为了生成伪随机数，需要设置一个种子
	//time.Now().Unix()返回从1970/1/1 0：0到现在的秒数
	rand.Seed(time.Now().Unix())
	//生成伪随机数
	num1 := rand.Intn(100) + 1
	fmt.Println(num1)
	//随机生成一个1-100的数字，直到生成99，记录一共需要使用多少次生成
	var num2 int
	var count int = 0
	for {
		num2 = rand.Intn(100) + 1
		count++
		fmt.Printf("%v\t", num2)
		if num2 == 99 {
			fmt.Print("Done!")
			break
		}
	}
	fmt.Println("count: ", count)

	//break label
label1:
	for i := 0; i < 4; i++ {
	label2:
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label2
			}
			if j == 1 {
				break label1
			}
			fmt.Println(j)
		}
	}

}
