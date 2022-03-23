package main

import "fmt"

func fbn(n int) []uint64 {
	fbn := make([]uint64, n)
	fbn[0] = 1
	fbn[1] = 1
	for i := 2; i < n; i++ {
		fbn[i] = fbn[i-1] + fbn[i-2]
	}
	return fbn

}
func main() {
	var num int
	fmt.Println("key a number:")
	fmt.Scanln(&num)
	fmt.Println(fbn(num))

}
