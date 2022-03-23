package main

import "fmt"

func main() {
	//map: key-value结构，字段、关联数组
	//方式1
	//声明:不会分配内存，无法使用。make初始化，分配内存后才能赋值使用
	var a map[string]string
	fmt.Println(a)
	//make分配内存
	a = make(map[string]string, 10)
	//添加和修改
	a["num1"] = "qinzhibao"
	fmt.Println(a)
	//方式2
	b := make(map[string]string)
	fmt.Println(b)
	//方式3
	c := map[string]string{"num1": "qin", "num2": "zhi"}
	//需要全部删除，可以遍历删除，或者make新空间
	delete(c, "num1") //删除

	fmt.Println(c)
	//应用
	students := make(map[string]map[string]string, 3)
	students["01"] = make(map[string]string)
	students["01"]["name"] = "tom"
	students["01"]["sex"] = "male"
	students["01"]["address"] = "bei jing"
	students["02"] = make(map[string]string)
	students["02"]["name"] = "jerry"
	students["02"]["sex"] = "male"
	students["02"]["address"] = "shang hai"
	fmt.Println(students)
	fmt.Println(students["02"]["name"])

	//查找
	value, isFind := students["02"]
	if isFind {
		fmt.Println(value)
	}

	//遍历：for-range
	for key, value := range students {
		fmt.Println(students[key], value)
	}
	for key1, value1 := range students {
		fmt.Printf("key1 = %v\n", key1)
		for key2, value2 := range value1 {
			fmt.Printf("key2 = %v,value2 = %v\n", key2, value2)
		}
	}
	//length
	fmt.Println(len(students))

}
