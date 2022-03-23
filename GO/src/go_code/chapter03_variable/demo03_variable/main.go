package main

import "fmt"

//声明定义全局变量
var n11 int = 100
var n22 = 200
var name2 = "jerry"

var (
	n33   = 300
	n44   = 400
	name3 = "jack"
)

func main() {
	//该案例演示golang一次性声明多个变量
	fmt.Println(n11, n22, n33, n44)
	//方式1
	var n1, n2, n3 int
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)

	//方式2
	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4 = ", n4, "name = ", name, "n5 = ", n5)
	//方式3
	n6, name1, n7 := 100, "tom", 888
	fmt.Println("n6 = ", n6, "name1 = ", name1, "n7 = ", n7)

	fmt.Println("name2 = ", name2, "name3 = ", name3)
}
