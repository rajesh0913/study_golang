package main

import "fmt"

func modify(map1 map[int]int) {
	map1[1] = 999
}

//定义struct
type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	//1. map是引用类型，遵守引用传递
	map1 := make(map[int]int)
	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	map1[4] = 4
	map1[5] = 5
	//2. 自动扩容
	map1[9] = 9
	fmt.Println("before: ", map1)
	modify(map1)
	fmt.Println("after: ", map1)
	//3. value经常使用struct管理复杂数据
	students := make(map[int]Stu)
	stu1 := Stu{Name: "tom", Age: 99, Address: "beijing"}
	stu2 := Stu{Name: "jerry", Age: 88, Address: "shanghai"}
	students[1] = stu1
	students[2] = stu2
	//fmt.Println(students)
	for index, value := range students {
		fmt.Printf("id: %v\t", index)
		fmt.Printf("name: %v\t", value.Name)
		fmt.Printf("age: %v\t\t", value.Age)
		fmt.Printf("address: %v\t", value.Address)
		fmt.Println()
	}
}
