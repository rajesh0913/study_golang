package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test.....")
	return true
}

//操作运算符
func main() {
	var n1 int = 0
	var n2 int = 8
	//关系运算符
	fmt.Println(n1 == n2)
	fmt.Println(n1 != n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 <= n2)
	flag := n1 > n2
	fmt.Printf("flag type:\t%T,\t\tflag = %v\n", flag, flag)
	//逻辑运算符
	//第一个条件为真就成立，不判断第二个条件
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}
	//在判断时，第一个条件满足时，||不会判断第二个条件
	i := 10
	if i > 9 || test() {
		fmt.Print("1:test ok!\n\n")
	} else {
		fmt.Print("1:test fail!\n\n")
	}
	if test() || i > 9 {
		fmt.Print("2:test ok!\n\n")
	} else {
		fmt.Print("2:test fail!\n\n")
	}
	//在判断时，第一个条件[不]满足时，&&不会判断第二个条件
	if i < 9 && test() {
		fmt.Print("3:test fail!\n\n")
	} else {
		fmt.Print("3:test ok!\n\n")
	}
	if test() && i > 9 {
		fmt.Print("4:test ok!\n\n")
	} else {
		fmt.Print("4:test fail!\n\n")
	}

}
