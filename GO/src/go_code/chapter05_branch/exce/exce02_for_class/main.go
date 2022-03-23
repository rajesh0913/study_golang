package main

import (
	"fmt"
)

func main() {
	studentNum := 0
	classNum := 0
	passcount := 0
	fmt.Println("输入班级个数：")
	fmt.Scanln(&classNum)
	fmt.Println("输入每个班学生个数：")
	fmt.Scanln(&studentNum)
	totalSum := 0.0
	for j := 1; j <= classNum; j++ {
		sum := 0.0
		for i := 1; i < studentNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班级第%d学生成绩:\t\n", j, i)
			fmt.Scanln(&score)
			if score >= 60 {
				passcount++
			}
			sum += score
		}

		fmt.Printf("%d班级平均分是：\t%v\n", j, sum/5)
		totalSum += sum
	}
	fmt.Printf("%v个班级总分是：\t%v,及格人数是：\t%v,平均分是：\t%v\n",
		classNum, totalSum, passcount, totalSum/float64(studentNum*classNum))
}
