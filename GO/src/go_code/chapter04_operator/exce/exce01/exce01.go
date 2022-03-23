package main

import (
	"fmt"
)

func main() {
	//97天放假，问多少个星期零多少天
	days_all := 97
	weeks := days_all / 7
	days_left := days_all % 7
	fmt.Println("weeks:\t", weeks, "\n", "days:\t", days_left)
	//定义一个变量保存华氏温度
	//公式：5/9*(华氏温度-100)，求出对应摄氏温度
	temp_h := 134.2
	temp_s := 5.0 / 9 * (temp_h - 100)
	fmt.Println("华氏温度：\t", temp_h, "\n", "摄氏温度：\t", temp_s)

}
