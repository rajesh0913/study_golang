package main

import (
	"fmt"
)

//defer():延时机制，函数执行之后及时释放资源
func sum(n1 int, n2 int) int {
	//执行到defer，会将其后的程序语句压入独立栈中
	//函数执行完毕之后，从defer栈中"先入后出"出栈
	//入栈时，相关变量也会拷贝入栈
	defer fmt.Println("n1 = ", n1) //3.n1 = 10
	defer fmt.Println("n2 = ", n2) //2.n2 = 20

	n1++ //n1=11
	n2++ //n2=21

	res := n1 + n2
	defer fmt.Println("res = ", res) //1.res = 32
	return res
}

func main() {

	fmt.Println("sum = ", sum(10, 20)) //4.sum = 32
	fmt.Println()

}
