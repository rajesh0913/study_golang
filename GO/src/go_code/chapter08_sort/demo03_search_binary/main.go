package main

import "fmt"

func binary_search(arr *[6]int, Left int, Right int, search_val int) {

	if Left > Right {
		fmt.Printf("no %v in the array!\n", search_val)
		return
	}

	middle := (Left + Right) / 2
	if (*arr)[middle] > search_val {
		binary_search(arr, Left, middle-1, search_val)
	} else if (*arr)[middle] < search_val {
		binary_search(arr, middle+1, Right, search_val)
	} else {
		fmt.Printf("find %v in %v address\n", search_val, middle)
	}
}

//前提是为有序数组
func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	var num int
	fmt.Println("send a number:")
	fmt.Scanln(&num)
	binary_search(&arr, 0, len(arr)-1, num)
}
