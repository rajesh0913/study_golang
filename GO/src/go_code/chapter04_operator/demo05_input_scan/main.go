package main

import (
	"fmt"
)

//从控制台获取用户信息：姓名、年龄、工资、是否通过考试
func main() {
	//方式1：fmt.scanln()
	var name string
	var years uint8
	var salary int
	var isPass bool
	fmt.Println("name:")
	fmt.Scanln(&name)
	fmt.Println("years:")
	fmt.Scanln(&years)
	fmt.Println("salary:")
	fmt.Scanln(&salary)
	fmt.Println("pass?:")
	fmt.Scanln(&isPass)

	fmt.Printf("name:\t%v\nyears:\t%v\nsalary:\t%v\npass:\t%v\n", name, years, salary, isPass)

	//方式2：fmt.scanf()
	fmt.Scanf("%s %d %d %t", &name, &years, &salary, &isPass)
	fmt.Printf("name:\t%v\nyears:\t%v\nsalary:\t%v\npass:\t%v\n", name, years, salary, isPass)

}
