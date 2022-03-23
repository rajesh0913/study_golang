package main

import "fmt"

type MethodUtils struct {
}

//print "*"
func (MethodUtils MethodUtils) Rectangle() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//print "*"
func (MethodUtils MethodUtils) Rectangle2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//area
func (MethodUtils MethodUtils) Area(width float64, len float64) (Area float64) {
	return width * len
}

//judge number
func (Method *MethodUtils) judgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

//print "*"
func (meth *MethodUtils) Rectangle3(m int, n int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}
func main() {
	var Method MethodUtils
	Method.Rectangle()
	fmt.Println()
	Method.Rectangle2(5, 4)
	fmt.Println()
	fmt.Println(Method.Area(5, 4))
	fmt.Println()
	(&Method).judgeNum(20)
	fmt.Println()
	(&Method).Rectangle3(10, 10, "Q")
}
