package main

import (
	"fmt"
)

//可变参数
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

}
