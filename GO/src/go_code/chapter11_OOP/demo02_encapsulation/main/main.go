package main

import (
	"fmt"
	"go_code/chapter11_OOP/demo02_encapsulation/encapsulation"
)

//封装：把抽象出的字段和字段的操作封装在一起，数据被保护在内存中，外部只有通过对应字段的操作才可以进行访问

func main() {
	p := encapsulation.NewPerson("qinzhibao")
	(*p).SetAge(25)
	(*p).SetSalary(15000)
	fmt.Println((*p).Name, (*p).GetAge())
}
