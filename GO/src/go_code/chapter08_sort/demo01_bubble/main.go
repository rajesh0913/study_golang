package main

import "fmt"

func Bubble_Sort(arr *[5]int) {
	fmt.Println("原数组： ", *arr)
	fmt.Println("数组长度：", len(*arr))
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			fmt.Printf("arr[%v] = %v\n", j, (*arr)[j])
			if (*arr)[i] < (*arr)[j] {
				temp := (*arr)[i]
				(*arr)[i] = (*arr)[j]
				(*arr)[j] = temp
			}
		}
		fmt.Println()
	}
	fmt.Println("排序后： ", *arr)
}
func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	Bubble_Sort(&arr)

}
