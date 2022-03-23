package main

// import "fmt"
// import "unsafe"
import (
	"fmt"
)

func main() {
	//演示字符类型使用
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1 = ", c1, "c2 = ", c2)
	//fomat print
	fmt.Printf("c1 = %c\t", c1)
	fmt.Printf("c2 = %c\n", c2)
	//overflow
	var c3 int = '北'
	fmt.Printf("c3 = %c ,c3 对应码值 = %d", c3, c3)
	fmt.Println("\a")

	var n1 = 10 + 'a'
	fmt.Printf("n1 = %c\tn1 = %d", n1, n1)

}
