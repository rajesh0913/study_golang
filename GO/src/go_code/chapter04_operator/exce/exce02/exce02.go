package main

import (
	"fmt"
)

//不使用中间变量，交换a b的值
func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)

}
