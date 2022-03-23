package main

import "fmt"

func main() {
	//1. 切片可以继续切片
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1[1:3]
	fmt.Println(slice2)
	//2. append()进行增加->创建底层新数组扩容，slice2重新指向新数组
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2)
	fmt.Printf("slice2 address = %p\n", &slice2[0])
	slice2 = append(slice2, slice2...)
	fmt.Println(slice2)
	fmt.Printf("slice2 new address = %p\n", &slice2[0])
	//3. 拷贝操作
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 10, 20)
	slice4[9] = 99
	copy(slice4, slice3) //数据类型为切片
	fmt.Println(slice4)
	//【copy不会扩展】
	slice5 := make([]int, 1)
	fmt.Printf("slice5 = %v\n", slice5)
	copy(slice5, slice3)
	fmt.Printf("copied slice5(no enough cap) = %v\n", slice5)
	copy(slice4, slice3)
	fmt.Printf("copied slice4(enough cap) = %v\n", slice4)

	//4. 切片是引用类型，执行引用传递
}
